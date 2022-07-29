package config

import (
	"github.com/spf13/viper"
)

var Instance *Config

type Config struct {
	Env        string `yaml:"Env"`        // 环境：prod、dev
	BaseUrl    string `yaml:"BaseUrl"`    // base url
	Port       string `yaml:"Port"`       // 端口
	LogFile    string `yaml:"LogFile"`    // 日志文件
	ShowSql    bool   `yaml:"ShowSql"`    // 是否显示日志
	StaticPath string `yaml:"StaticPath"` // 静态文件目录

	// 数据库配置
	DB struct {
		Url          string `yaml:"Url"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`

	// Github
	Github struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"Github"`

	// OSChina
	OSChina struct {
		ClientID     string `yaml:"ClientID"`
		ClientSecret string `yaml:"ClientSecret"`
	} `yaml:"OSChina"`

	// QQ登录
	QQConnect struct {
		AppId  string `yaml:"AppId"`
		AppKey string `yaml:"AppKey"`
	} `yaml:"QQConnect"`

	// 阿里云oss配置
	Uploader struct {
		Enable    string `yaml:"Enable"`
		AliyunOss struct {
			Host          string `yaml:"Host"`
			Bucket        string `yaml:"Bucket"`
			Endpoint      string `yaml:"Endpoint"`
			AccessId      string `yaml:"AccessId"`
			AccessSecret  string `yaml:"AccessSecret"`
			StyleSplitter string `yaml:"StyleSplitter"`
			StyleAvatar   string `yaml:"StyleAvatar"`
			StylePreview  string `yaml:"StylePreview"`
			StyleSmall    string `yaml:"StyleSmall"`
			StyleDetail   string `yaml:"StyleDetail"`
		} `yaml:"AliyunOss"`
		Local struct {
			Host string `yaml:"Host"`
			Path string `yaml:"Path"`
		} `yaml:"Local"`
	} `yaml:"Uploader"`

	// 百度ai
	BaiduAi struct {
		ApiKey    string `yaml:"ApiKey"`
		SecretKey string `yaml:"SecretKey"`
	} `yaml:"BaiduAi"`

	// 百度SEO相关配置
	// 文档：https://ziyuan.baidu.com/college/courseinfo?id=267&page=2#h2_article_title14
	BaiduSEO struct {
		Site  string `yaml:"Site"`
		Token string `yaml:"Token"`
	} `yaml:"BaiduSEO"`

	// 神马搜索SEO相关
	// 文档：https://zhanzhang.sm.cn/open/mip
	SmSEO struct {
		Site     string `yaml:"Site"`
		UserName string `yaml:"UserName"`
		Token    string `yaml:"Token"`
	} `yaml:"SmSEO"`

	// smtp
	Smtp struct {
		Host     string `yaml:"Host"`
		Port     string `yaml:"Port"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		SSL      bool   `yaml:"SSL"`
	} `yaml:"Smtp"`

	// es
	Es struct {
		Url   string `yaml:"Url"`
		Index string `yaml:"Index"`
	} `yaml:"Es"`
}

func Init(filename string) *Config {
	// 项目中原生的读取配置的方法
	//Instance = &Config{}
	//if yamlFile, err := ioutil.ReadFile(filename); err != nil {
	//	logrus.Error(err)
	//} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
	//	logrus.Error(err)
	//}

	/*
			利用viper读配置有几个优点：
			1.支持多种格式的配置文件， JSON/TOML/YAML/HCL/envfile/Java properties等，
		      原项目中的 yaml.Unmarshal 只支持解析 yaml(java常用)；
			2.可以设置监听配置文件的修改，修改时自动加载新的配置；这是最大的魅力，它使得重启服务器才能使配置生效的日子一去不复返！
			3.从环境变量、命令行选项和io。reader中读取配置；
			4.从远程配置系统中读取和监听修改，如 etcd/Consul；
			5.代码逻辑中显示设置键值。
	*/
	conf := viper.New()

	conf.AddConfigPath(filename) // 设置读取的文件路径
	conf.SetConfigName("bbs-go") // 设置读取的文件名
	conf.SetConfigType("yaml")
	if err := conf.ReadInConfig(); err != nil {
		panic(err)
	}
	// 将配置文件内容映射到 struct 中
	if err := conf.Unmarshal(&Instance); err != nil { // 注意这里 参数 必须是指针类型！
		panic(err)
	}

	return Instance
}
