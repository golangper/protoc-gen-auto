package plugin

import (
	"fmt"
	"strings"

	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/golangper/protoc-gen-rorm/options"
)

type RormPlugin struct {
	*generator.Generator
	imports map[generator.GoPackageName]generator.GoImportPath

	file *generator.FileDescriptor
}

type Api struct {
	method   string
	path     string
	funcName string
}

// Name identifies the plugin
func (p *RormPlugin) Name() string {
	return "rorm"
}

func (p *RormPlugin) Init(g *generator.Generator) {
	p.Generator = g
	p.imports = make(map[generator.GoPackageName]generator.GoImportPath)
}

func (p *RormPlugin) Generate(file *generator.FileDescriptor) {
	p.file = file
	p.imports["json"] = "encoding/json"
	p.P(`var _ = json.Marshal`)
	p.imports["grpc"] = "google.golang.org/grpc"

	for _, svc := range file.GetService() {
		var redisType int64 = 0
		var xormType int64 = 0
		var useUid bool = false
		var useNsq bool = false
		var _ = useUid
		value, err := proto.GetExtension(svc.Options, options.E_RedisType)
		if err != nil || value == nil {
			// fmt.Println("===", err)
			redisType = 0
		} else {
			redisType = *(value.(*int64))
		}
		value, err = proto.GetExtension(svc.Options, options.E_XormType)
		if err != nil || value == nil {
			xormType = 0
		} else {
			xormType = *(value.(*int64))
		}

		if xormType > 0 {
			p.imports["xorm"] = "github.com/go-xorm/xorm"
		} else {
			delete(p.imports, "xorm")
		}
		if redisType > 0 {
			p.imports["redis"] = "github.com/go-redis/redis"
			//p.imports["strconv"] = "strconv"
		} else {
			delete(p.imports, "redis")
		}

		useNsq = proto.GetBoolExtension(svc.Options, options.E_UseNsq, false)

		if useNsq {
			p.imports["nsqpool"] = "github.com/qgymje/nsqpool"
		} else {
			delete(p.imports, "nsqpool")
		}
		//route := GetApiRouteExtension(svc.Options)

		gin := proto.GetBoolExtension(svc.Options, options.E_GinHandler, false)

		if gin {
			p.imports["http"] = "net/http"
			p.imports["gin"] = "github.com/gin-gonic/gin"
			// p.imports["binding"] = "github.com/gin-gonic/gin/binding"
		} else {
			delete(p.imports, "http")
			delete(p.imports, "gin")
		}
		p.imports["log"] = "log"
		p.imports["context"] = "golang.org/x/net/context"
		p.imports["roundrobin"] = "google.golang.org/grpc/balancer/roundrobin"
		grpcSvcName := "_" + generator.CamelCase(svc.GetName()) + "Imp"
		impName := generator.CamelCase(svc.GetName()) + "Imp"
		//grpc impl struct
		p.P(`type `, grpcSvcName, ` struct {`)
		p.In()
		if xormType == 1 {
			p.P(`db *xorm.Engine`)
		} else if xormType == 2 {
			p.P(`db *xorm.EngineGroup`)
		}
		if redisType == 1 {
			p.P(`redis *redis.Client`)
		} else if redisType == 2 {
			p.P(`redis *redis.ClusterClient`)
		}
		if useNsq {
			p.P(`nsq *pool.Pool`)
		}
		p.Out()
		p.P(`}`)

		apilist := make([]*Api, 0)
		//grpc method impl
		for _, method := range svc.GetMethod() {
			mname := generator.CamelCase(method.GetName())

			inputType := generator.CamelCase(method.GetInputType())
			outputType := generator.CamelCase(method.GetOutputType())

			uid := GetUidExtension(method.Options)
			opts := GetOptsExtension(method.Options)
			strs := strings.Split(inputType, ".")
			inputMsg := p.AllFiles().GetMessage(strs[1], strs[2])
			//inputMsg := p.file.GetMessage(GetMessageName(method.GetInputType()))
			if inputMsg == nil {
				fmt.Println("inputType: ", inputType, "not find")
				return
			}
			if uid != nil {
				err := CheckUidSeed(inputMsg, uid)
				if err != nil {
					fmt.Println(err.Error())
					return
				}
			}
			var in, out string
			if file.GetPackage() == strs[1] {
				in = generator.CamelCase(strs[2])
			} else {
				in = strs[1] + "." + generator.CamelCase(strs[2])
			}
			ots := strings.Split(outputType, ".")
			if file.GetPackage() == ots[1] {
				out = generator.CamelCase(ots[2])
			} else {
				out = ots[1] + "." + generator.CamelCase(ots[2])
			}
			// in := inputType[1:]
			// out := outputType[1:]
			p.P(``)
			p.P(`func (s *`, grpcSvcName, `) `, mname, `(c context.Context, in *`, in, `) (*`, out, `, error) {`)
			p.In()
			p.P(`var err error`)
			p.outAndValid(in, out)
			if uid != nil {
				p.newUid(uid)
			}
			if opts != nil {
				p.dealMethods(opts, method.GetInputType(), method.GetOutputType(), mname)
			}

			p.P(`return out, nil`)
			p.Out()
			p.P(`}`)

		}
		p.P(``)
		p.P(`type `, impName, ` struct {`)
		p.In()
		p.P(grpcSvcName)
		p.Out()
		p.P(`}`)

		prm := ""
		if xormType == 1 {
			if prm != "" {
				prm += ", "
			}
			prm += `db *xorm.Engine`
		} else if xormType == 2 {
			if prm != "" {
				prm += ", "
			}
			prm += `db *xorm.EngineGroup`
		}
		if redisType == 1 {
			if prm != "" {
				prm += ", "
			}
			prm += `redis *redis.Client`
		} else if redisType == 2 {
			if prm != "" {
				prm += ", "
			}
			prm += `redis *redis.ClusterClient`
		}
		if useNsq {
			if prm != "" {
				prm += ", "
			}
			prm += `nsq *pool.Pool`
		}
		if prm != "" {
			prm += ", "
		}
		p.P(``)
		p.P(`func New`, impName, `(`, prm, `) `, impName, ` {`)
		p.In()
		p.P(`res := `, impName, `{}`)
		if xormType > 0 {
			p.P(`res.db = db`)
		}
		if redisType > 0 {
			p.P(`res.redis = redis`)
		}
		if useNsq {
			p.P(`res.nsq = nsq`)
		}
		p.P(`return res`)
		p.Out()
		p.P(`}`)
		p.P(``)

		for _, m := range svc.GetMethod() {
			mname := generator.CamelCase(m.GetName())
			//inputType := generator.CamelCase(m.GetInputType())
			api := GetApiExtension(m.Options)
			if api != nil {
				myapi := &Api{method: api.Method, path: api.Path, funcName: mname + "Handler"}
				apilist = append(apilist, myapi)
			}
			inputType := generator.CamelCase(m.GetInputType())
			strs := strings.Split(inputType, ".")
			var in string
			if file.GetPackage() == strs[1] {
				in = generator.CamelCase(strs[2])
			} else {
				in = strs[1] + "." + generator.CamelCase(strs[2])
			}
			p.P(``)
			p.P(`func (s *`, impName, `) `, mname, `Handler(c *gin.Context) {`)
			p.In()

			p.P(`prm := &`, in, `{}`)
			p.P(`var err error`)

			p.P(`err = c.ShouldBind(prm)`)
			p.P(`if err != nil {`)
			p.In()
			p.P(`log.Println("`, mname+"Handler[c.ShouldBind] :", `", err.Error())`)
			p.P(`c.String(http.StatusBadRequest, err.Error())`)
			p.P(`return`)
			p.Out()
			p.P(`}`)

			// p.P(`if err = prm.Validate(); err != nil {`)
			// p.In()
			// p.P(`log.Println("`,mname + "Handler[prm.Validate] :",`", err.Error())`)
			// p.P(`c.String(http.StatusBadRequest, err.Error())`)
			// p.P(`return`)
			// p.Out()
			// p.P(`}`)

			p.P(`res, err := s.`, mname, `(context.Background(), prm)`)
			p.P(`if err != nil {`)
			p.In()
			p.P(`log.Println("`, mname+"Handler[s."+mname+"] :", `", err.Error())`)
			p.P(`c.String(http.StatusServiceUnavailable, err.Error())`)
			p.P(`return`)
			p.Out()
			p.P(`}`)
			p.P(`r,_:=json.Marshal(res)`)
			p.P(`c.String(http.StatusOK,string(r))`)
			p.Out()
			p.P(`}`)
		}
		// p.GenerateImports(file)
		p.P(``)
		p.P(`func (s *`, impName, `) InitApi(g *gin.RouterGroup) {`)
		p.In()
		for _, l := range apilist {
			if l.method == "post" || l.method == "POST" || l.method == "Post" {
				p.P(`g.POST("`, l.path, `", s.`, l.funcName, `)`)
			} else if l.method == "get" || l.method == "GET" || l.method == "Get" {
				p.P(`g.GET("`, l.path, `", s.`, l.funcName, `)`)
			} else {
				fmt.Println("not not support the method", l.method)
			}
		}
		p.Out()
		p.P(`}`)

		p.P(``)
		p.P(`func (s *`, impName, `) NewBalancerClient(target string) `, generator.CamelCase(svc.GetName()), `Client {`)
		p.In()
		p.P(`conn, err := grpc.Dial(target,grpc.WithInsecure(),grpc.WithBalancerName(roundrobin.Name))`)
		p.P(`if err != nil {`)
		p.In()
		p.P(`log.Fatalln(err)`)
		p.Out()
		p.P(`}`)
		p.P(`return New`, generator.CamelCase(svc.GetName()), `Client(conn)`)
		p.Out()
		p.P(`}`)

	}
}

func (p *RormPlugin) outAndValid(in, out string) {
	if in == out {
		p.P(`out := in`)
	} else {
		p.P(`out := &`, out, `{}`)
	}
	p.P(`err = in.Validate()`)
	p.P(`if err != nil {`)
	p.In()
	p.P(`return out, err`)
	p.Out()
	p.P(`}`)
}

func (p *RormPlugin) newUid(uid *options.UidOptions) {
	p.imports["snowflake"] = "github.com/fainted/snowflake"
	strs := strings.Split(uid.Seed, ".")
	f := strs[len(strs)-1]
	p.P(`_s := in.`, generator.CamelCase(f), ` % 256`)
	p.P(`_worker, err := snowflake.NewChannelWorker(_s)`)
	p.P(`if err != nil {`)
	p.In()
	p.P("return out, err")
	p.Out()
	p.P(`}`)
	p.P(uid.Name, ` , _ := _worker.Next()`)
	p.P(`var _ = `, uid.Name)
}

func (p *RormPlugin) dealMethods(opts *options.RormOptions, in, out, mname string) {
	err := p.dealMethod(opts, false, false, in, out, mname)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (p *RormPlugin) dealMethod(opt *options.RormOptions, end bool, els bool, in, out, mname string) error {
	//fmt.Println(opt)
	var err error
	if opt.GetParam() == "" && opt.GetSqlxTran() == nil {
		return fmt.Errorf("param can not bu null")
	}
	param := strings.Replace(opt.GetParam(), "\n", "", -1)

	//str := strings.Replace(param, `'`, `"`, -1)

	strArry := strings.Split(param, ";")
	str1 := `"` + strings.TrimSpace(strArry[0]) + `"`;
	str2 := " "
	for _, s := range strArry[1:] {
		str2 += ","
		str2 += CamelField(strings.TrimSpace(s))
	}
	//str = strings.Replace(param, `;`, `,`, -1)
	//var tp descriptor.FieldDescriptorProto_Type
	tp, lb, sl, err := p.getVarType(opt.GetTarget(), in, out)
	if err != nil {
		return err
	}
	switch opt.GetMethod() {
	case "xorm.Exec":
		if lb == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Errorf("xorm.Exec's target can not be repeated ")
		}
		p.P(`_, err = s.db.Exec(`, str1, str2, `)`)
		p.dealErrBool(opt, tp, mname+"[s.db.Exec]:")
	case "xorm.SQLGet":
		if lb == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Errorf("xorm.SQLGet's target can not be repeated ")
		}
		if tp == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			if sl != "" {
				p.P(`for _, obj := range `, CamelField(sl), `{`)
				p.In()
				s := strings.Split(CamelField(opt.GetTarget()), ".")
				p.P(`_, err = s.db.SQL(`, str1, str2, `).Get(obj.`, s[len(s)-1], `)`)
				p.Out()
				p.P(`}`)
			} else {
				p.P(`_, err = s.db.SQL(`, str1, str2, `).Get(`, CamelField(opt.GetTarget()), `)`)
			}

		} else {
			if sl != "" {
				p.P(`for _, obj := range `, CamelField(sl), `{`)
				p.In()
				s := strings.Split(CamelField(opt.GetTarget()), ".")
				p.P(`_, err = s.db.SQL(`, str1, str2, `).Get(&obj.`, s[len(s)-1], `)`)
				p.Out()
				p.P(`}`)
			} else {
				p.P(`_, err = s.db.SQL(`, str1, str2, `).Get(&`, CamelField(opt.GetTarget()), `)`)
			}
		}
		if opt.Failure == nil {
			p.dealErrReturn(mname + "[s.db.SQL-Get]:")
		}
	case "xorm.SQLFind":
		if lb != descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Errorf("xorm.SQLFind's target must be repeated ")
		}
		if sl != "" {
			p.P(`for _, obj := range `, CamelField(sl), `{`)
			p.In()
			s := strings.Split(CamelField(opt.GetTarget()), ".")
			p.P(`err = s.db.SQL(`, str1, str2, `).Find(&obj.`, s[len(s)-1], `)`)
			p.Out()
			p.P(`}`)
		} else {
			p.P(`err = s.db.SQL(`, str1, str2, `).Find(&`, CamelField(opt.GetTarget()), `)`)
		}
		if opt.Failure == nil {
			p.dealErrReturn(mname + "[s.db.SQL-Find]:")
		}
	case "redis.Get":
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		t := strings.Split(CamelField(opt.GetTarget()), ".")
		n := t[len(t)-1]
		switch tp {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			p.P(`rds`, n, `, err := s.redis.Get(`, key, `).Bytes()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
			}

			p.P(`err = `, CamelField(opt.GetTarget()), `.Unmarshal(rds`, n, `)`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[" + CamelField(opt.GetTarget()) + "Unmarshal]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_STRING:
			p.P(CamelField(opt.GetTarget()), `, err := s.redis.Get(`, key, `).String()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_INT64:
			p.P(CamelField(opt.Target), `, err := s.redis.Get(`, key, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_UINT64, descriptor.FieldDescriptorProto_TYPE_FIXED64:
			p.P(`rds`, n, `, err := s.redis.Get(`, key, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
				p.P(CamelField(opt.Target), ` = uint64(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = uint64(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_INT32:
			p.P(`rds`, n, `, err := s.redis.Get(`, key, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
				p.P(CamelField(opt.Target), ` = int32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = int32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_FIXED32:
			p.P(`rds`, n, `, err := s.redis.Get(`, key, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
				p.P(CamelField(opt.Target), ` = uint32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = uint32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			p.P(CamelField(opt.Target), `, err := s.redis.Get(`, key, `).Float64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			p.P(`rds`, n, `, err := s.redis.Get(`, key, `).Float64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.Get]:")
				p.P(CamelField(opt.Target), ` = float32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = float32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		default:
			return fmt.Errorf("redis.Get's target type can not support ")
		}

	case "redis.Set":
		if len(strArry) < 3 {
			return fmt.Errorf("redis.Set's param must have 2 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		tp1, _, _, err := p.getVarType(strArry[1], in, out)
		if err != nil {
			return err
		}
		if !checkStrIsNum(strings.TrimSpace(strArry[2])) {
			return fmt.Errorf("redis.Set's the thired param must be int ")
		}
		t := strings.Split(CamelField(opt.GetTarget()), ".")
		n := t[len(t)-1]
		p.imports["time"] = "time"
		if tp1 == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			p.P(`set`, n, `, err := `, CamelField(strArry[1]), `.Marshal()`)
			p.dealErrReturn(mname + "[Marshal]:")
			p.P(`err = s.redis.Set(`, key, `,`, `set`, n, `,int64(time.Duration(`, strArry[2], `) * time.Second)).Err()`)
		} else if tp1 == descriptor.FieldDescriptorProto_TYPE_BOOL {
			return fmt.Errorf("redis.Set's target can not be bool ")
		} else {
			p.P(`err = s.redis.Set(`, key, `,`, CamelField(strArry[1]), `,int64(time.Duration(`, strArry[2], `) * time.Second)).Err()`)
		}
		p.dealErrBool(opt, tp, mname+"[s.redis.Set]")
	case "redis.Del":
		param := ""
		for _, st := range strArry {
			key, err := p.getString(strings.TrimSpace(st), in, out)
			if err != nil {
				return err
			}
			if param != "" {
				param += ","
			}
			param += key
		}
		p.P(`err = s.redis.Del(`, param, `).Err()`)
		p.dealErrBool(opt, tp, mname+"[s.redis.Del]")
	case "redis.IncrByX":
		if len(strArry) < 2 {
			return fmt.Errorf("redis.IncrByX's param must have 2 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		tp1, _, _, err := p.getVarType(opt.GetTarget(), in, out)
		if err != nil {
			return err
		}
		switch tp1 {
		case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT32:
			p.P(`err = s.redis.IncrBy(`, key, `, int64(`, CamelField(strArry[1]), `)).Err()`)
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE, descriptor.FieldDescriptorProto_TYPE_FLOAT:
			p.P(`err = s.redis.IncrByFloat(`, key, `, float64(`, CamelField(strArry[1]), `)).Err()`)
		default:
			return fmt.Errorf("redis.IncrByX's The second param can be int32  int64 float32 float64")
		}
		p.dealErrBool(opt, tp, mname+"[s.redis.IncrBy|IncrByFloat|IncrByX]")
	case "redis.DecrBy":
		if len(strArry) < 2 {
			return fmt.Errorf("redis.DecrBy's param must have 2 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		tp1, _, _, err := p.getVarType(opt.GetTarget(), in, out)
		if err != nil {
			return err
		}
		switch tp1 {
		case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT32:
			p.P(`_dnum,  err := s.redis.Get(`, key, `).Int64()`)
			p.dealErrReturn(mname + "[s.redis.Get]:")
			p.P(`if int(_dnum) < int(`, CamelField(strArry[1]), `){`)
			p.In()
			p.P(`return out, fmt.Errorf("Inventory shortage")`)
			p.Out()
			p.P(`}`)
			p.P(`err = s.redis.IncrBy(`, key, `, int64(`, CamelField(strArry[1]), `)).Err()`)
		default:
			return fmt.Errorf("redis.IncrByX's The second param can be int32  int64")
		}
		p.dealErrBool(opt, tp, mname+"[s.redis.IncrBy]")
	case "redis.Expire":
		if len(strArry) < 2 {
			return fmt.Errorf("redis.Expire's param must have 2 ")
		}
		if !checkStrIsInt(strings.TrimSpace(strArry[1])) {
			return fmt.Errorf("redis.Expire: The second param must be int num")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		p.imports["time"] = "time"
		p.P(`err = s.redis.Expire(`, key, `, int64(time.Duration(`, strArry[1], `) * time.Second)).Err()`)
		p.dealErrBool(opt, tp, mname+"[s.redis.Expire]")
	case "redis.HGet":
		if len(strArry) < 2 {
			return fmt.Errorf("redis.HGet's param must have 2 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}

		field, err := p.getString(strArry[1], in, out)
		if err != nil {
			return err
		}
		t := strings.Split(CamelField(opt.GetTarget()), ".")
		n := t[len(t)-1]
		switch tp {
		case descriptor.FieldDescriptorProto_TYPE_MESSAGE:
			p.P(`rds`, n, `, err := s.redis.HGet(`, key, `,`, field, `).Bytes()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
			}

			p.P(`err = `, CamelField(opt.GetTarget()), `.Unmarshal(rds`, n, `)`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[Unmarshal]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_STRING:
			p.P(CamelField(opt.GetTarget()), `, err := s.redis.HGet(`, key, `,`, field, `).String()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_INT64:
			p.P(CamelField(opt.GetTarget()), `, err := s.redis.HGet(`, key, `,`, field, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_UINT64, descriptor.FieldDescriptorProto_TYPE_FIXED64:
			p.P(`rds`, n, `, err := s.redis.HGet(`, key, `,`, field, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
				p.P(CamelField(opt.Target), ` = uint64(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = uint64(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_INT32:
			p.P(`rds`, n, `, err := s.redis.HGet(`, key, `,`, field, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
				p.P(CamelField(opt.Target), ` = int32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = int32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_FIXED32:
			p.P(`rds`, n, `, err := s.redis.HGet(`, key, `,`, field, `).Int64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
				p.P(CamelField(opt.Target), ` = uint32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = uint32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
			p.P(CamelField(opt.GetTarget()), `, err := s.redis.HGet(`, key, `,`, field, `).Float64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
			}
		case descriptor.FieldDescriptorProto_TYPE_FLOAT:
			p.P(`rds`, n, `, err := s.redis.HGet(`, key, `,`, field, `).Float64()`)
			if opt.Failure == nil {
				p.dealErrReturn(mname + "[s.redis.HGet]:")
				p.P(CamelField(opt.Target), ` = float32(`, `rds`, n, `)`)
			} else {
				p.P(`if err == nil {`)
				p.In()
				p.P(CamelField(opt.Target), ` = float32(`, `rds`, n, `)`)
				p.Out()
				p.P(`}`)
			}
		default:
			return fmt.Errorf("redis.Get's target can not be repeated ")
		}
	case "redis.HSet":
		if len(strArry) < 3 {
			return fmt.Errorf("redis.HSet's param must have 3 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		_, lb2, _, err := p.getVarType(strArry[1], in, out)
		if err != nil {
			return err
		}
		if lb2 == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Errorf("redis.HSet field param can not be repeated ")
		}
		field, err := p.getString(strArry[1], in, out)
		if err != nil {
			return err
		}

		tp1, lb1, _, err := p.getVarType(strArry[2], in, out)
		if err != nil {
			return err
		}
		if lb1 == descriptor.FieldDescriptorProto_LABEL_REPEATED {
			return fmt.Errorf("redis.HSet value param can not be repeated ")
		}

		t := strings.Split(CamelField(strArry[2]), ".")
		n := t[len(t)-1]
		if tp1 == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			p.P(`set`, n, `, err := `, CamelField(strArry[2]), `.Marshal()`)
			p.dealErrReturn(mname + "[Marshal]:")
			p.P(`err = s.redis.HSet(`, key, `,`, field, `, set`, n, `).Err()`)
		} else if tp1 == descriptor.FieldDescriptorProto_TYPE_BOOL {
			return fmt.Errorf("redis.Set's target can not be bool ")
		} else {
			p.P(`err = s.redis.HSet(`, key, `,`, field, `,`, CamelField(strArry[2]), `).Err()`)
		}
		p.dealErrBool(opt, tp, mname+"[s.redis.HSet]")
	case "redis.HDel":
		if len(strArry) < 2 {
			return fmt.Errorf("redis.HDel's param must have 2 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		field, err := p.getString(strArry[1], in, out)
		p.P(`err := s.redis.HDel(`, key, `,`, field, `).Err()`)
		p.dealErrBool(opt, tp, mname+"[s.redis.HDel]")
	case "redis.HincrByX":
		if len(strArry) < 3 {
			return fmt.Errorf("redis.HincrByX's param must have 3 ")
		}
		key, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		field, err := p.getString(strArry[1], in, out)
		if err != nil {
			return err
		}
		tp2, _, _, err := p.getVarType(strArry[2], in, out)
		if err != nil {
			return err
		}
		switch tp2 {
		case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_INT32:
			p.P(`err = s.redis.HIncrBy(`, key, `,`, field, `, int64(`, CamelField(strArry[2]), `)).Err()`)
		case descriptor.FieldDescriptorProto_TYPE_DOUBLE, descriptor.FieldDescriptorProto_TYPE_FLOAT:
			p.P(`err = s.redis.HIncrByFloat(`, key, `,`, field, `, float64(`, CamelField(strArry[2]), `)).Err()`)
		default:
			return fmt.Errorf("redis.HincrByX's The second param can be int32  int64 float32 float64")
		}
		p.dealErrBool(opt, tp, mname+"[s.redis.HIncrBy|HIncrByFloat]")
	case "nsq.Producer":
		if len(strArry) < 2 {
			return fmt.Errorf("nsq.Producer's param must have 3 ")
		}
		topic, err := p.getString(str1, in, out)
		if err != nil {
			return err
		}
		tp1, _, _, err := p.getVarType(strArry[1], in, out)
		if err != nil {
			return err
		}
		if tp1 == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			p.P(`pdc, err := (*s.nsq).Get()`)
			p.P(`defer pdc.Close()`)
			p.dealErrReturn(mname + "[(*s.nsq).Get()]:")
			p.P(`_nsqData, err := `, CamelField(strArry[1]), `.Marshal()`)
			p.dealErrReturn(mname + "[Marshal]:")
			p.P(`err = pdc.Publish(`, topic, ` _nsqData)`)
			p.dealErrReturn(mname + "[pdc.Publish]:")
		} else {
			return fmt.Errorf("nsq.Producer: the second param must be message ")
		}
	default:
		if opt.GetSqlxTran() != nil && len(opt.GetSqlxTran()) > 0 {
			p.P(`tx := s.db.NewSession()`)
			p.P(`err = tx.Begin()`)
			p.dealErrReturn(mname + "[tx.Begin]:")
			for _, o := range opt.GetSqlxTran() {
				//str := strings.Replace(o.GetParam(), `'`, `"`, -1)
				strArry := strings.Split(o.GetParam(), ";")
				str1 :=`"` + strings.TrimSpace(strArry[0]) + `"`
				str2 := " "
				for _, s := range strArry[1:] {
					str2 += ","
					str2 += CamelField(strings.TrimSpace(s))
				}
				_, lb1, _, err := p.getVarType(o.GetSlice(), in, out)
				if err != nil {
					fmt.Println(err)
					return err
				}
				if o.Method == "xorm.Exec" {
					if lb1 == descriptor.FieldDescriptorProto_LABEL_REPEATED {
						p.P(`for _, obj := range `, CamelField(o.GetSlice()), `{`)
						p.In()
						p.P(`_, err = s.db.Exec(`, str1, str2, `)`)
						p.P(`if err != nil {`)
						p.In()
						p.P(`tx.Rollback()`)
						p.P(`log.Println("`, mname+"[s.db.Exec] :", `", err.Error())`)
						p.P(`return out, err`)
						p.Out()
						p.P(`}`)
						p.Out()
						p.P(`}`)

					} else {
						p.P(`_, err = s.db.Exec(`, str1, str2, `)`)
						p.P(`if err != nil {`)
						p.In()
						p.P(`tx.Rollback()`)
						p.P(`log.Println("`, mname+"[s.db.Exec] :", `", err.Error())`)
						p.P(`return out, err`)
						p.Out()
						p.P(`}`)
					}

				} else {
					fmt.Println("Does not support functions: %s", opt.GetMethod())
					err = fmt.Errorf("Does not support functions: %s", opt.GetMethod())
				}
			}
			p.P(`tx.Commit()`)
		} else if opt.GetMzset() != nil {

		} else {
			fmt.Println("Does not support functions: %s", opt.GetMethod())
			err = fmt.Errorf("Does not support functions: %s", opt.GetMethod())
		}
	}
	if end {
		p.Out()
		if els {
			p.P(`} else {`)
			p.In()
		} else {
			p.P(`}`)
		}
	}
	if opt.Success != nil && opt.Failure != nil {
		p.P(`if err != nil {`)
		p.In()
		p.P(`log.Println("`, mname+"[...] :", `", err.Error())`)
		err = p.dealMethod(opt.Failure, true, true, in, out, mname)
		if err == nil {
			err = p.dealMethod(opt.Success, true, false, in, out, mname)
		}
	} else if opt.Failure != nil {
		p.P(`if err != nil {`)
		p.In()
		p.P(`log.Println("`, mname+"[...] :", `", err.Error())`)
		err = p.dealMethod(opt.Failure, true, false, in, out, mname)
	} else if opt.Success != nil {
		err = p.dealMethod(opt.Success, false, false, in, out, mname)
	}
	return err
}
func (p *RormPlugin) dealErrReturn(mname string) {
	p.P(`if err != nil {`)
	p.In()
	p.P(`log.Println("`, mname, `", err.Error())`)
	p.P(`return out, err`)
	p.Out()
	p.P(`}`)
}
func (p *RormPlugin) dealErrBool(opt *options.RormOptions, tp descriptor.FieldDescriptorProto_Type, mname string) {
	if opt.Failure == nil {
		p.P(`if err != nil {`)
		p.In()
		p.P(`log.Println("`, mname, `", err.Error())`)
		p.P(`return out, err`)
		p.Out()
		p.P(`}`)
		if descriptor.FieldDescriptorProto_TYPE_BOOL == tp {
			p.P(CamelField(opt.Target), ` = true`)
		}
	} else {
		if descriptor.FieldDescriptorProto_TYPE_BOOL == tp {
			p.P(`if err == nil {`)
			p.In()
			p.P(CamelField(opt.Target), ` = true`)
			p.Out()
			p.P(`}`)
		}
	}
}
func (p *RormPlugin) getString(str, in, out string) (string, error) {
	s := strings.Replace(str, " ", "", -1)
	ss := strings.Split(s, "+")
	res := ""
	for _, st := range ss {
		if !strings.Contains(st, `in.`) {
			if res != "" {
				res += " + "
			}
			res += st
		} else {
			vars := strings.Split(st, ".")
			if vars[0] != "in" {
				return "", fmt.Errorf("param %s is not valide", st)
			}
			// msg := p.file.GetMessage(in)
			strs := strings.Split(in, ".")
			msg := p.AllFiles().GetMessage(strs[1], strs[2])
			var tp descriptor.FieldDescriptorProto_Type
			for _, f := range vars[1:] {
				fd := msg.GetFieldDescriptor(strings.TrimSpace(f))
				tp = fd.GetType()
				if tp == descriptor.FieldDescriptorProto_TYPE_MESSAGE {

					if msg == nil {
						return "", fmt.Errorf("can not find message %s in this file", fd.GetTypeName())
					}
					strs = strings.Split(fd.GetTypeName(), ".")
					msg = p.AllFiles().GetMessage(strs[1], strs[2])
				}
			}

			switch tp {
			case descriptor.FieldDescriptorProto_TYPE_STRING:
				if res != "" {
					res += " + "
				}
				res += CamelField(st)
			case descriptor.FieldDescriptorProto_TYPE_INT64, descriptor.FieldDescriptorProto_TYPE_UINT64,
				descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_FIXED64,
				descriptor.FieldDescriptorProto_TYPE_FIXED32:
				if res != "" {
					res += " + "
				}
				p.imports["strconv"] = "strconv"
				res += "strconv.Itoa(int(" + CamelField(st) + "))"
			case descriptor.FieldDescriptorProto_TYPE_FLOAT:
				if res != "" {
					res += " + "
				}
				p.imports["strconv"] = "strconv"
				res += "strconv.FormatFloat(" + CamelField(st) + ",'f',-1,32)"
			case descriptor.FieldDescriptorProto_TYPE_DOUBLE:
				if res != "" {
					res += " + "
				}
				p.imports["strconv"] = "strconv"
				res += "strconv.FormatFloat(" + CamelField(st) + ",'f',-1,64)"
			default:
				return "", fmt.Errorf("field %s can not to string", st)
			}
		}
	}
	return res, nil
}

func (p *RormPlugin) getVarType(st string, in, out string) (descriptor.FieldDescriptorProto_Type, descriptor.FieldDescriptorProto_Label, string, error) {
	if st == "" {
		return 0, 0, "", nil
	}
	vars := strings.Split(st, ".")
	var msg *descriptor.DescriptorProto
	if vars[0] == "in" {
		strs := strings.Split(in, ".")
		msg = p.AllFiles().GetMessage(strs[1], strs[2])
	} else if vars[0] == "out" {
		strs := strings.Split(out, ".")
		msg = p.AllFiles().GetMessage(strs[1], strs[2])
	} else {
		return 0, 0, "", fmt.Errorf("target must start with  'in' or 'out' ")
	}
	if len(vars) == 1 {
		return descriptor.FieldDescriptorProto_TYPE_MESSAGE, descriptor.FieldDescriptorProto_LABEL_OPTIONAL, "", nil
	}
	var tp descriptor.FieldDescriptorProto_Type
	var lb descriptor.FieldDescriptorProto_Label
	var sl string
	for i, f := range vars {
		if i == 0 {
			sl += f
			continue
		}
		if i < len(vars)-1 {
			sl += "."
			sl += strings.TrimSpace(f)
		}

		fd := msg.GetFieldDescriptor(strings.TrimSpace(f))
		if fd == nil {
			return 0, 0, "", fmt.Errorf("can not find field %s in this file", strings.TrimSpace(f))
		}
		tp = fd.GetType()
		lb = fd.GetLabel()
		if i == len(vars)-2 {
			if lb != descriptor.FieldDescriptorProto_LABEL_REPEATED {
				sl = ""
			}
		}
		if len(vars) < 3 {
			sl = ""
		}
		if tp == descriptor.FieldDescriptorProto_TYPE_MESSAGE {
			if msg == nil {
				return 0, 0, "", fmt.Errorf("can not find message %s in this file", fd.GetTypeName())
			}
			strs := strings.Split(fd.GetTypeName(), ".")
			msg = p.AllFiles().GetMessage(strs[1], strs[2])
			// msg = p.file.GetMessage(GetMessageName(fd.GetTypeName()))
		}
	}
	return tp, lb, sl, nil
}
