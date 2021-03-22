#礼品码接口
- 开发语言：golang  *v1.16*
- 后端框架：gin  *v1.6.3*
- 数据库：redis-v8  *v8.7.1*
##1.基本介绍
###1.1项目介绍
- 本地测试地址：
  - [创建礼品码](http://localhost:8080/createCode)
  - [验证礼品码](http://localhost:8080/giftVerify/)
  
- 开发者：lyh
- 当前版本：1.0
###1.2使用说明
>- 安装gin
   >>"github.com/gin-gonic/gin"
>- 安装redis-v8
>>"github.com/go-redis/redis/v8"

- 创建礼品接口 
  - 接收前端gift参数  
    1.gift数据格式：
    >{"count":Count,"gold":Gold,"diamond":Diamond,"prop":Prop}
    
    2.gift参数信息：

    | 参数 | 含义 | 备注
    | :------| :------|:------|
    | count | 礼品码领取次数 | 负数代表不限领取次数，默认为0
    | gold | 礼包金币数量    | 默认为0
    | diamond | 礼包钻石数量 | 默认为0
    | prop | 礼包道具数量    | 默认为0

  - 接口返回礼品码code参数  
    1.code参数格式
    >8位礼品码 例：AE8S2DZX
- 验证礼品接口
  - 接收前端礼品码参数：code
  - 返回礼品码发放结果
    >发放成功：返回gift参数  
     发放失败：返回错误信息

##2.项目说明
###2.1目录结构
- log (保存日志文件)
- src (项目源码)
  - config (项目配置文件)
  - controller (应用层源码)
  - dao (数据库源码)
  - entity (实体源码)
  - log (日志系统配置)
  - router (路由配置)
  - service (逻辑层源码)
  - setting (项目配置)
  - util (工具包)
- testUtil (单元测试文件)

###2.2 应用层
- **CreatCode函数**  
1. 调用util.bind()函数绑定前端gift参数  
2. 调用service.CreateGift()函数获取礼品码
3. 返回礼品码创建结果
  >成功返回礼品码code / 失败返回错误信息

- **VerifyCode函数**
1. 读取cookie验证用户登陆信息
2. 读取code参数，判定礼品码是否为空
3. 调用service.VerifyGift()获取礼品信息
> 礼品码合法返回礼品内容，执行步骤4 / 礼品码非法返回""字符串，返回错误信息
4. 调用service.AddGift()增加奖励
5. 返回礼品内容

###2.3 逻辑层
- **CreateGift(gift string)函数**
1. 获取礼品信息
2. 调用util.Encode(code int64)获取生成的8位礼品码
3. 调用dao.Set()方法存储礼品码以及对应的礼品信息
>成功返回json / 失败返回""字符串
4. 返回礼品码code

- **VerifyGift(code string,gift entity.AddGift)函数**
1. 获取礼品码code
2. 查询数据库code是否存在
>code不存在返回错误信息
3. 检查礼品信息中领取次数是否到达上限  
> count > 0 未达上限，更新count--  
> count == 0 礼品领取礼品领取到达上限，调用dao.Del删除code  
> count < 0 礼品领取不限次数
4. 检查用户是否重复领取  
>查询数据库是否有该用户信息领取信息
5. 返回领取结果json字符串

- **AddGift(uid string,gift entity.AddGift)函数**
1. 获取uid和礼品信息，执行增加礼品逻辑  
2. 保存用户礼品领取信息
3. 返回增加的礼品信息json字符串

###2.4 dao层
- **Set(code string,gift string)函数**
1. 调用Redis.Set() 存储礼品码信息
>key = GIFTCODE + code  
> value = gift

- **Get(code string, clear bool)函数**
1. 调用Redis.Get(code) 查询数据库
2. clear判断获取后是否删除key

- **Exist(code string)函数**
1. 查询key是否存在

- **Del(code string)函数**
1. 删除key=code

###2.5 单元测试
>礼品结构：  
> type AddGift struct {  
&ensp;&ensp;&ensp;&ensp;Count   int `json:"count"`       //礼品数量 负数为无限领取  
&ensp;&ensp;&ensp;&ensp;Gold    int `json:"gold"`        //增加金币数量  
&ensp;&ensp;&ensp;&ensp;Diamond int `json:"diamond"`     //增加钻石数量  
&ensp;&ensp;&ensp;&ensp;Prop    int `json:"prop"`        //增加道具数量  
}  
> 测试输入：  
> var g AddGift  
&ensp;&ensp;&ensp;&ensp;g.Count = 5  
&ensp;&ensp;&ensp;&ensp;g.Gold = 5  
&ensp;&ensp;&ensp;&ensp;g.Diamond = 5  
&ensp;&ensp;&ensp;&ensp;g.Prop = 5  
- **TestEncode()**  
  
  Encode函数测试
  礼品id编码为8位礼品码，
  测试礼品id = 1-10代码输出，预期结果为不同的8为字符串  
- **TestBind()**  
  Bind绑定JSON字符串测试
  访问localhost:8080/TestBind地址，传递json，控制台打印Json字符串  
- **TestNewSetting()**   
  1.NewSetting函数测试单元  
  输出配置文件所有信息  

  2.ReadSection函数测试单元  
  输出Reids配置文件所有信息
  
- **TestVerify()**  
  Verify函数单元测试  
  控制台打印返回结果
  > 成功返回礼品内容/失败打印错误
- **TestCreate()**  
  Create函数测试类  
  输入定礼品内容，打印生成的8位礼品码字符串
  
- **TestAddGift()**  
  输入礼品内容，测试AddGift函数,控制台打印返回增加礼品结果
  
- **TestCreate()**  
  Create函数测试单元  
  访问localhost:8080/TestCreatCode  
  发送post请求  
  >json：  
  {  
  &ensp;&ensp;&ensp;&ensp;"count": 5,  
  &ensp;&ensp;&ensp;&ensp;"gold": 5,  
  &ensp;&ensp;&ensp;&ensp;"diamond": 5,  
  &ensp;&ensp;&ensp;&ensp;"prop": 5  
  }  
  打印返回创建的礼品码  
  
- **TestVerifyCode()**  
  VerifyCode函数测试单元  
  访问localhost:8080/TestVerifyCode  
  >参数:  
  &ensp;&ensp;&ensp;&ensp;cookie：userName=121  
  &ensp;&ensp;&ensp;&ensp;query：code=8A8S2DZX  

  返回打印领取结果  