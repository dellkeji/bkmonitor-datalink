// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package config

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/viper"
	"golang.org/x/exp/slices"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

var (
	// FilePath path of logger
	FilePath = "./bmw.yaml"
	// EnvKeyPrefix env prefix
	EnvKeyPrefix = "bmw"

	// BrokerRedisMode redis mode
	BrokerRedisMode string
	// BrokerRedisSentinelMasterName broker redis mater name
	BrokerRedisSentinelMasterName string
	// BrokerRedisSentinelAddress redis address
	BrokerRedisSentinelAddress []string
	// BrokerRedisSentinelPassword password of broker redis
	BrokerRedisSentinelPassword string
	// BrokerRedisStandaloneHost host of standalone broker redis
	BrokerRedisStandaloneHost string
	// BrokerRedisStandalonePort port of standalone broker redis
	BrokerRedisStandalonePort int
	// BrokerRedisStandalonePassword password of standalone broker redis
	BrokerRedisStandalonePassword string
	// BrokerRedisDatabase db of broker redis
	BrokerRedisDatabase int
	// BrokerRedisDialTimeout broker redis dial timeout
	BrokerRedisDialTimeout time.Duration
	// BrokerRedisReadTimeout broker redis dial timeout
	BrokerRedisReadTimeout time.Duration
	// BrokerRedisMetricPublishKey key for metrics publish/subscribe
	BrokerRedisMetricPublishKey string

	// StorageRedisMode mode of storage redis
	StorageRedisMode string
	// StorageRedisSentinelMasterName master name of storage redis
	StorageRedisSentinelMasterName string
	// StorageRedisSentinelAddress address of storage redis
	StorageRedisSentinelAddress []string
	// StorageRedisSentinelPassword password of storage redis
	StorageRedisSentinelPassword string
	// StorageRedisStandaloneHost host of storage redis
	StorageRedisStandaloneHost string
	// StorageRedisStandalonePort port of storage redis
	StorageRedisStandalonePort int
	// StorageRedisStandalonePassword password of storage redis
	StorageRedisStandalonePassword string
	// StorageRedisDatabase db of storage redis
	StorageRedisDatabase int
	// StorageRedisDialTimeout storage redis dial timeout
	StorageRedisDialTimeout time.Duration
	// StorageRedisReadTimeout storage redis read timeout
	StorageRedisReadTimeout time.Duration
	// StorageRedisKeyPrefix storage prefix
	StorageRedisKeyPrefix string

	// StorageDependentRedisMode dependent redis mode
	StorageDependentRedisMode string
	// StorageDependentRedisSentinelMasterName dependent redis master name
	StorageDependentRedisSentinelMasterName string
	// StorageDependentRedisSentinelAddress dependent redis address
	StorageDependentRedisSentinelAddress []string
	//StorageDependentRedisSentinelPassword dependent redis password
	StorageDependentRedisSentinelPassword string
	// StorageDependentRedisStandaloneHost dependent redis host
	StorageDependentRedisStandaloneHost string
	// StorageDependentRedisStandalonePort dependent redis(standalone) port
	StorageDependentRedisStandalonePort int
	// StorageDependentRedisStandalonePassword dependent redis(standalone) password
	StorageDependentRedisStandalonePassword string
	// StorageDependentRedisDatabase dependent redis db
	StorageDependentRedisDatabase int
	// StorageDependentRedisDialTimeout dependent redis dial timeout
	StorageDependentRedisDialTimeout time.Duration
	// StorageDependentRedisReadTimeout dependent redis read timeout
	StorageDependentRedisReadTimeout time.Duration

	// StorageConsulPathPrefix prefix of consul
	StorageConsulPathPrefix string
	// StorageConsulSrvName consul server name
	StorageConsulSrvName string
	// StorageConsulAddress consul address
	StorageConsulAddress string
	// StorageConsulPort consul port
	StorageConsulPort int
	// StorageConsulAddr consul address
	StorageConsulAddr string
	// StorageConsulTag tag of consul
	StorageConsulTag []string
	// StorageConsulTll consul ttl
	StorageConsulTll string

	// StorageMysqlHost mysql host
	StorageMysqlHost string
	// StorageMysqlPort mysql port
	StorageMysqlPort int
	// StorageMysqlUser mysql user
	StorageMysqlUser string
	// StorageMysqlPassword mysql password
	StorageMysqlPassword string
	// StorageMysqlDbName mysql db
	StorageMysqlDbName string
	// StorageMysqlCharset mysql charset
	StorageMysqlCharset string
	// StorageMysqlMaxIdleConnections mysql max idle
	StorageMysqlMaxIdleConnections int
	// StorageMysqlMaxOpenConnections mysql max open size
	StorageMysqlMaxOpenConnections int
	// StorageMysqlDebug enabled mysql debug
	StorageMysqlDebug bool

	// StorageEsUpdateTaskRetainInvalidAlias whether retain invalid alias
	StorageEsUpdateTaskRetainInvalidAlias bool

	// StorageBboltDefaultPath bbolt default path
	StorageBboltDefaultPath string
	// StorageBboltDefaultBucketName bbolt default bucket name
	StorageBboltDefaultBucketName string
	// StorageBboltDefaultSync bbolt default sync
	StorageBboltDefaultSync bool

	// WorkerQueues worker listen queue(only valid in worker process)
	WorkerQueues []string
	// WorkerConcurrency concurrency of worker task
	WorkerConcurrency int
	// WorkerHealthCheckInterval interval of worker report health status
	WorkerHealthCheckInterval time.Duration
	// WorkerHealthCheckInfoDuration cache duration of worker info
	WorkerHealthCheckInfoDuration time.Duration
	// WorkerDaemonTaskMaintainerInterval check interval of task maintainer
	WorkerDaemonTaskMaintainerInterval time.Duration
	// WorkerDaemonTaskRetryTolerateCount max retry of task
	WorkerDaemonTaskRetryTolerateCount int
	// WorkerDaemonTaskRetryTolerateInterval retry interval of failed task
	WorkerDaemonTaskRetryTolerateInterval time.Duration
	// WorkerDaemonTaskRetryIntolerantFactor retry duration factor of failed task
	WorkerDaemonTaskRetryIntolerantFactor int

	// SchedulerTaskWatchChanSize Listen for the maximum number of concurrent tasks in the broker queue
	SchedulerTaskWatchChanSize int
	// SchedulerDaemonTaskNumeratorInterval interval of scheduler numerator
	SchedulerDaemonTaskNumeratorInterval time.Duration
	// SchedulerDaemonTaskWorkerWatcherInterval interval of scheduler worker watcher
	SchedulerDaemonTaskWorkerWatcherInterval time.Duration
	// SchedulerDaemonTaskTaskWatcherInterval interval of scheduler task watcher
	SchedulerDaemonTaskTaskWatcherInterval time.Duration

	// GinMode http mode
	GinMode string
	// TaskListenHost http listen host
	TaskListenHost string
	// TaskListenPort http listen port
	TaskListenPort int
	// ControllerListenHost http listen host
	ControllerListenHost string
	// TaskListenPort http listen port
	ControllerListenPort int
	// TaskListenHost http listen host
	WorkerListenHost string
	// TaskListenPort http listen port
	WorkerListenPort int

	// AesKey project aes key
	AesKey string

	// BkApiEnabled enabled bk-apigw
	BkApiEnabled bool
	// BkApiUrl bk-apigw host
	BkApiUrl string
	// BkApiStage bk-apigw stage
	BkApiStage string
	// BkApiAppCode bk-apigw app code
	BkApiAppCode string
	// BkApiAppSecret bk-apigw app secret
	BkApiAppSecret string
	// BkApiBcsApiGatewayDomain bk-apigw bcs domain
	BkApiBcsApiGatewayDomain string
	// BkApiBcsApiGatewayToken bk-apigw bcs token
	BkApiBcsApiGatewayToken string

	// GoroutineLimit max size of task goroutine
	GoroutineLimit map[string]string

	// TestStorageMysqlHost test-mysql host
	TestStorageMysqlHost string
	// TestStorageMysqlPort test-mysql port
	TestStorageMysqlPort int
	// TestStorageMysqlUser test-mysql user
	TestStorageMysqlUser string
	// TestStorageMysqlPassword test-mysql password
	TestStorageMysqlPassword string
	// TestStorageMysqlDbName test-mysql db name
	TestStorageMysqlDbName string
)

func initVariables() {

	/* Storage Redis 配置 */
	StorageRedisMode = GetValue("store.redis.mode", "standalone")
	StorageRedisSentinelMasterName = GetValue("store.redis.sentinel.masterName", "")
	StorageRedisSentinelAddress = GetValue("store.redis.sentinel.address", []string{"127.0.0.1"})
	StorageRedisSentinelPassword = GetValue("store.redis.sentinel.password", "")
	StorageRedisStandaloneHost = GetValue("store.redis.standalone.host", "127.0.0.1")
	StorageRedisStandalonePort = GetValue("store.redis.standalone.port", 6379)
	StorageRedisStandalonePassword = GetValue("store.redis.standalone.password", "")
	StorageRedisDatabase = GetValue("store.redis.db", 0)
	StorageRedisDialTimeout = GetValue("store.redis.dialTimeout", 10*time.Second, viper.GetDuration)
	StorageRedisReadTimeout = GetValue("store.redis.readTimeout", 10*time.Second, viper.GetDuration)
	StorageRedisKeyPrefix = GetValue("store.redis.keyPrefix", "bmw")

	/* Storage DependentRedis 配置 */
	StorageDependentRedisMode = GetValue("store.dependentRedis.mode", "standalone")
	StorageDependentRedisSentinelMasterName = GetValue("store.dependentRedis.sentinel.masterName", "")
	StorageDependentRedisSentinelAddress = GetValue("store.dependentRedis.sentinel.address", []string{"127.0.0.1"})
	StorageDependentRedisSentinelPassword = GetValue("store.dependentRedis.sentinel.password", "")
	StorageDependentRedisStandaloneHost = GetValue("store.dependentRedis.standalone.host", "127.0.0.1")
	StorageDependentRedisStandalonePort = GetValue("store.dependentRedis.standalone.port", 6379)
	StorageDependentRedisStandalonePassword = GetValue("store.dependentRedis.standalone.password", "")
	StorageDependentRedisDatabase = GetValue("store.dependentRedis.db", 0)
	StorageDependentRedisDialTimeout = GetValue("store.dependentRedis.dialTimeout", 10*time.Second, viper.GetDuration)
	StorageDependentRedisReadTimeout = GetValue("store.dependentRedis.readTimeout", 10*time.Second, viper.GetDuration)

	/* Storage Consul配置 */
	StorageConsulPathPrefix = GetValue("store.consul.pathPrefix", "bk_bkmonitorv3_enterprise_production")
	StorageConsulSrvName = GetValue("store.consul.srvName", "bmw")
	StorageConsulAddress = GetValue("store.consul.address", "127.0.0.1:8500")
	StorageConsulPort = GetValue("store.consul.port", 8500)
	StorageConsulAddr = GetValue("store.consul.addr", "http://127.0.0.1:8500")
	StorageConsulTag = GetValue("store.consul.tag", []string{"bmw"})
	StorageConsulTll = GetValue("store.consul.ttl", "")

	/* Storage Mysql配置 */
	StorageMysqlHost = GetValue("store.mysql.host", "127.0.0.1")
	StorageMysqlPort = GetValue("store.mysql.port", 3306)
	StorageMysqlUser = GetValue("store.mysql.user", "root")
	StorageMysqlPassword = GetValue("store.mysql.password", "")
	StorageMysqlDbName = GetValue("store.mysql.dbName", "")
	StorageMysqlCharset = GetValue("store.mysql.charset", "utf8")
	StorageMysqlMaxIdleConnections = GetValue("store.mysql.maxIdleConnections", 10)
	StorageMysqlMaxOpenConnections = GetValue("store.mysql.maxOpenConnections", 100)
	StorageMysqlDebug = GetValue("store.mysql.debug", false)

	StorageEsUpdateTaskRetainInvalidAlias = GetValue("store.es.esRetainInvalidAlias", false)

	StorageBboltDefaultPath = GetValue("store.bbolt.defaultPath", "bolt.db")
	StorageBboltDefaultBucketName = GetValue("store.bbolt.defaultBuckName", "spaceBucket")
	StorageBboltDefaultSync = GetValue("store.bbolt.defaultSync", false)

	/*
		Worker配置 ----- START
	*/
	// WorkerQueues worker进行监听的队列名称列表 在worker启动时可以通过--queues="x1,x2"指定 不指定默认使用default队列
	WorkerQueues = GetValue("worker.queues", []string{"default"})
	// WorkerConcurrency worker并发数量 0为使用CPU核数
	WorkerConcurrency = GetValue("worker.concurrency", 0)
	// WorkerHealthCheckInterval worker心跳上报时间间隔
	WorkerHealthCheckInterval = GetValue("worker.healthCheck.interval", 3*time.Second, viper.GetDuration)
	// WorkerHealthCheckInfoDuration worker心跳上报缓存过期时间
	WorkerHealthCheckInfoDuration = GetValue("worker.healthCheck.duration", 5*time.Second, viper.GetDuration)
	// WorkerDaemonTaskMaintainerInterval worker常驻任务检测任务是否正常运行的间隔
	WorkerDaemonTaskMaintainerInterval = GetValue(
		"worker.daemonTask.maintainer.interval", 1*time.Second, viper.GetDuration,
	)
	// WorkerDaemonTaskRetryTolerateCount worker常驻任务配置，当任务重试超过指定数量仍然失败时，下次重试间隔就不断动态增长
	WorkerDaemonTaskRetryTolerateCount = GetValue("worker.daemonTask.maintainer.tolerateCount", 60)
	// WorkerDaemonTaskRetryTolerateInterval worker常驻任务当任务执行失败并且重试次数未超过 WorkerDaemonTaskRetryTolerateCount 时
	// 下次重试时间间隔
	WorkerDaemonTaskRetryTolerateInterval = GetValue(
		"worker.daemonTask.maintainer.tolerateInterval", 10*time.Second, viper.GetDuration,
	)
	// WorkerDaemonTaskRetryIntolerantFactor worker常驻任务当任务重试次数超过 WorkerDaemonTaskRetryTolerateCount 时
	// 下次重试按照Nx倍数增长 设置倍数因子
	WorkerDaemonTaskRetryIntolerantFactor = GetValue("worker.daemonTask.maintainer.intolerantFactor", 2)
	/*
		Worker配置 ----- END
	*/

	/*
		Scheduler常驻任务配置 ----- START
	*/
	// SchedulerTaskWatchChanSize 调度器监听定时任务最大并发数量
	SchedulerTaskWatchChanSize = GetValue("scheduler.watcher.chanSize", 10)
	// SchedulerDaemonTaskNumeratorInterval 定时检测当前常驻任务分派是否正确的时间间隔(默认每60秒检测一次)
	SchedulerDaemonTaskNumeratorInterval = GetValue(
		"scheduler.daemonTask.numerator.interval", 60*time.Second, viper.GetDuration,
	)
	// SchedulerDaemonTaskWorkerWatcherInterval 常驻任务功能监听worker队列变化的间隔
	SchedulerDaemonTaskWorkerWatcherInterval = GetValue(
		"scheduler.daemonTask.watcher.workerWatchInterval", 1*time.Second, viper.GetDuration,
	)
	// SchedulerDaemonTaskTaskWatcherInterval 常驻任务功能监听task队列变化的间隔
	SchedulerDaemonTaskTaskWatcherInterval = GetValue(
		"scheduler.daemonTask.watcher.taskWatchInterval", 1*time.Second, viper.GetDuration,
	)
	/*
		Scheduler常驻任务配置 ----- END
	*/

	GinMode = GetValue("service.mode", "release")
	TaskListenHost = GetValue("service.task.listen", "127.0.0.1")
	TaskListenPort = GetValue("service.task.port", 10211)
	ControllerListenHost = GetValue("service.controller.listen", "127.0.0.1")
	ControllerListenPort = GetValue("service.controller.port", 10212)
	WorkerListenHost = GetValue("service.worker.listen", "127.0.0.1")
	WorkerListenPort = GetValue("service.worker.port", 10213)

	AesKey = GetValue("aes.key", "")

	BkApiEnabled = GetValue("taskConfig.common.bkapi.enabled", false)
	BkApiUrl = GetValue("taskConfig.common.bkapi.host", "http://127.0.0.1")
	BkApiStage = GetValue("taskConfig.common.bkapi.stage", "stag")
	BkApiAppCode = GetValue("taskConfig.common.bkapi.appCode", "appCode")
	BkApiAppSecret = GetValue("taskConfig.common.bkapi.appSecret", "appSecret")
	BkApiBcsApiGatewayDomain = GetValue("taskConfig.common.bkapi.bcsApiGatewayDomain", "")
	BkApiBcsApiGatewayToken = GetValue("taskConfig.common.bkapi.bcsApiGatewayToken", "")

	GoroutineLimit = GetValue("taskConfig.common.goroutineLimit", map[string]string{}, viper.GetStringMapString)

	TestStorageMysqlHost = GetValue("test.store.mysql.host", "127.0.0.1")
	TestStorageMysqlPort = GetValue("test.store.mysql.port", 3306)
	TestStorageMysqlUser = GetValue("test.store.mysql.user", "root")
	TestStorageMysqlPassword = GetValue("test.store.mysql.password", "")
	TestStorageMysqlDbName = GetValue("test.store.mysql.dbName", "")
}

var (
	keys []string
)

// 结构化配置
type BMWConfig struct {
	Service   ServiceConfig         `mapstructure:"service"`
	Broker    BrokerConfig          `mapstructure:"broker"`
	Store     StoreConfig           `mapstructure:"store"`
	Log       LogConfig             `mapstructure:"log"`
	Aes       AesConfig             `mapstructure:"aes"`
	Worker    WorkerModuleConfig    `mapstructure:"workerConfig"`
	Task      TaskModuleConfig      `mapstructure:"taskConfig"`
	Scheduler SchedulerModuleConfig `mapstructure:"schedulerConfig"`
}

// http service config
type ServiceConfig struct {
	Mode       string                  `mapstructure:"mode"`
	Task       TaskServiceConfig       `mapstructure:"task"`
	Controller ControllerServiceConfig `mapstructure:"controller"`
	Worker     WorkerServiceConfig     `mapstructure:"worker"`
}

// broker config
type BrokerConfig struct {
	RedisConfig RedisBrokerConfig `mapstructure:"redis"`
}

// store config
type StoreConfig struct {
	RedisConfig          RedisStoreConfig          `mapstructure:"redis"`
	DependentRedisConfig RedisDependentStoreConfig `mapstructure:"dependentRedis"`
	MysqlConfig          MysqlStoreConfig          `mapstructure:"mysql"`
	ConsulConfig         ConsulStoreConfig         `mapstructure:"consul"`
	EsConfig             EsStoreConfig             `mapstructure:"es"`
}

// log config
type LogConfig struct {
	EnableStdout bool   `mapstructure:"enableStdout"`
	Level        string `mapstructure:"level"`
	Path         string `mapstructure:"path"`
	MaxSize      int    `mapstructure:"maxSize"`
	MaxAge       int    `mapstructure:"maxAge"`
	MaxBackups   int    `mapstructure:"maxBackups"`
}

// aes config
type AesConfig struct {
	key string `mapstructure:"key"`
}

// worker module config
type WorkerModuleConfig struct {
	Concurrency      int                     `mapstructure:"concurrency"`
	Queues           []string                `mapstructure:"queues"`
	HealthCheck      WorkerHealthCheckConfig `mapstructure:"healthCheck"`
	DaemonTaskConfig WorkerDaemonTaskConfig  `mapstructure:"daemonTask"`
}

// task modul config
type TaskModuleConfig struct {
	Common          TaskModuleCommonConfig          `mapstructure:"common"`
	Metadata        TaskModuleMetadataConfig        `mapstructure:"metadata"`
	ApmPreCalculate TaskModuleApmPreCalculateConfig `mapstructure:"apmPreCalculate"`
}

// NOTE: 防止有特殊配置，先单独出来
// task config
type TaskServiceConfig struct {
	Listen string `mapstructure:"listen"`
	Port   int    `mapstructure:"port"`
}

// controller config
type ControllerServiceConfig struct {
	Listen string `mapstructure:"listen"`
	Port   int    `mapstructure:"port"`
}

// worker config
type WorkerServiceConfig struct {
	Listen string `mapstructure:"listen"`
	Port   int    `mapstructure:"port"`
}

// redis broker config
type RedisBrokerConfig struct {
	Mode        string                      `mapstructure:"mode"`
	DB          int                         `mapstructure:"db"`
	DialTimeout time.Duration               `mapstructure:"dialTimeout"`
	ReadTimeout time.Duration               `mapstructure:"readTimeout"`
	Standalone  StandaloneRedisBrokerConfig `mapstructure:"standalone"`
	Sentinel    SentinelRedisBrokerConfig   `mapstructure:"sentinel"`
}

// standalone broker redis
type StandaloneRedisBrokerConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// sentinel broker redis
type SentinelRedisBrokerConfig struct {
	MasterName string   `mapstructure:"masterName"`
	Address    []string `mapstructure:"address"`
	Password   string   `mapstructure:"password"`
}

// redis storage config
type RedisStoreConfig struct {
	Mode        string                     `mapstructure:"mode"`
	DB          int                        `mapstructure:"db"`
	DialTimeout time.Duration              `mapstructure:"dialTimeout"`
	ReadTimeout time.Duration              `mapstructure:"readTimeout"`
	Standalone  StandaloneRedisStoreConfig `mapstructure:"standalone"`
	Sentinel    SentinelRedisStoreConfig   `mapstructure:"sentinel"`
	KeyPrefix   string                     `mapstructure:"keyPrefix"`
}

type StandaloneRedisStoreConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// sentinel broker redis
type SentinelRedisStoreConfig struct {
	MasterName string   `mapstructure:"masterName"`
	Address    []string `mapstructure:"address"`
	Password   string   `mapstructure:"password"`
}

type RedisDependentStoreConfig struct {
	Mode        string                         `mapstructure:"mode"`
	DB          string                         `mapstructure:"db"`
	DialTimeout string                         `mapstructure:"dialTimeout"`
	ReadTimeout string                         `mapstructure:"readTimeout"`
	Standalone  StandaloneRedisDependentConfig `mapstructure:"standalone"`
	Sentinel    SentinelRedisDependentConfig   `mapstructure:"sentinel"`
}

// standalone broker redis
type StandaloneRedisDependentConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	password string `mapstructure:"password"`
}

// sentinel broker redis
type SentinelRedisDependentConfig struct {
	MasterName string   `mapstructure:"masterName"`
	Address    []string `mapstructure:"address"`
	Password   string   `mapstructure:"password"`
}

// mysql config
type MysqlStoreConfig struct {
	Debug              bool   `mapstructure:"debug"`
	Host               string `mapstructure:"host"`
	Port               int    `mapstructure:"port"`
	User               string `mapstructure:"user"`
	Password           string `mapstructure:"password"`
	DbName             string `mapstructure:"dbName"`
	Charset            string `mapstructure:"charset"`
	MaxIdleConnections int    `mapstructure:"maxIdleConnections"`
	MaxOpenConnections int    `mapstructure:"maxOpenConnections"`
}

// consul config
type ConsulStoreConfig struct {
	PathPrefix string   `mapstructure:"pathPrefix"`
	SrvName    string   `mapstructure:"srvName"`
	Address    string   `mapstructure:"address"`
	Port       int      `mapstructure:"port"`
	Addr       string   `mapstructure:"addr"`
	Tag        []string `mapstructure:"tag"`
	Ttl        string   `mapstructure:"ttl"`
}

// es config
type EsStoreConfig struct {
	EsRetainInvalidAlias bool `mapstructure:"esRetainInvalidAlias"`
}

// Worker health check config
type WorkerHealthCheckConfig struct {
	Interval time.Duration `mapstructure:"interval"`
	Duration time.Duration `mapstructure:"duration"`
}

// worker daemontask config
type WorkerDaemonTaskConfig struct {
	Maintainer DaemonTaskMaintainerConfig `mapstructure:"maintainer"`
}

// daemontask config
type DaemonTaskMaintainerConfig struct {
	Interval         time.Duration `mapstructure:"interval"`
	TolerateCount    int           `mapstructure:"tolerateCount"`
	TolerateInterval time.Duration `mapstructure:"tolerateInterval"`
	IntolerantFactor int           `mapstructure:"intolerantFactor"`
}

// task module common config
type TaskModuleCommonConfig struct {
	GoroutineLimit map[string]int        `mapstructure:"goroutineLimit"`
	BkApi          TaskModuleBkApiConfig `mapstructure:"bkapi"`
}

// task module bkapi
type TaskModuleBkApiConfig struct {
	Enabled             bool   `mapstructure:"enabled"`
	Host                string `mapstructure:"host"`
	Stage               string `mapstructure:"stage"`
	AppCode             string `mapstructure:"appCode"`
	AppSecret           string `mapstructure:"appSecret"`
	BcsApiGatewayDomain string `mapstructure:"bcsApiGatewayDomain"`
	BcsApiGatewayToken  string `mapstructure:"bcsApiGatewayToken"`
}

// task module metadata config
type TaskModuleMetadataConfig struct {
	MetricDimension TaskModuleMetadataMetricDimensionConfig `mapstructure:"metricDimension"`
	BcsConfig       TaskModuleBcsConfig                     `mapstructure:"bcs"`
}

// metadata metric config
type TaskModuleMetadataMetricDimensionConfig struct {
	MetricKeyPrefix             string `mapstructure:"metricKeyPrefix"`
	MetricDimensionKeyPrefix    string `mapstructure:"metricDimensionKeyPrefix"`
	MaxMetricsFetchStep         int    `mapstructure:"maxMetricsFetchStep"`
	TimeSeriesMetricExpiredDays int    `mapstructure:"timeSeriesMetricExpiredDays"`
}

// task module bcs config
type TaskModuleBcsConfig struct {
	EnableBcsGray         bool   `mapstructure:"enableBcsGray"`
	ClusterBkEnvLabel     string `mapstructure:"clusterBkEnvLabel"`
	KafkaStorageClusterId int    `mapstructure:"kafkaStorageClusterId"`
}

// apm pre calculate config
type TaskModuleApmPreCalculateConfig struct {
	Notifier map[string]int              `mapstructure:"notifier"`
	Window   ApmPreCalculateWindowConfig `mapstructure:"ApmPreCalculateWindowConfig"`
}

// apm pre calculate window config
type ApmPreCalculateWindowConfig struct {
	MaxSize                 int                               `mapstructure:"maxSize"`
	ExpireInterval          time.Duration                     `mapstructure:"expireInterval"`
	MaxDuration             time.Duration                     `mapstructure:"maxDuration"`
	ExpireIntervalIncrement int                               `mapstructure:"expireIntervalIncrement"`
	NoDataMaxDuration       time.Duration                     `mapstructure:"noDataMaxDuration"`
	Distributive            ApmPreCalculateDistributiveConfig `mapstructure:"distributive"`
	Processor               map[string]int                    `mapstructure:"processor"`
	Storage                 ApmPreCalculateStorageConfig      `mapstructure:"storage"`
}

// apm pre calculate distributive config
type ApmPreCalculateDistributiveConfig struct {
	SubSize                     int           `mapstructure:"subSize"`
	WatchExpireInterval         time.Duration `mapstructure:"watchExpireInterval"`
	ConcurrentCount             int           `mapstructure:"concurrentCount"`
	ConcurrentExpirationMaximum int           `mapstructure:"concurrentExpirationMaximum"`
}

// apm pre calculate storage config
type ApmPreCalculateStorageConfig struct {
	SaveRequestBufferSize int                        `mapstructure:"saveRequestBufferSize"`
	WorkerCount           int                        `mapstructure:"workerCount"`
	SaveHoldMaxCount      int                        `mapstructure:"saveHoldMaxCount"`
	SaveHoldMaxDuration   time.Duration              `mapstructure:"saveHoldMaxDuration"`
	Bloom                 ApmPreCalculateBloomConfig `mapstructure:"bloom"`
}

type ApmPreCalculateBloomConfig struct {
	FpRate        float64                  `mapstructure:"fpRate"`
	Normal        map[string]time.Duration `mapstructure:"normal"`
	NormalOverlap map[string]time.Duration `mapstructure:"normalOverlap"`
	LayersBloom   map[string]int           `mapstructure:"layersBloom"`
	DecreaseBloom map[string]int           `mapstructure:"decreaseBloom"`
}

// scheduler config
type SchedulerModuleConfig struct {
	Watcher    map[string]int                  `mapstructure:"watcher"`
	DaemonTask SchedulerModuleDaemonTaskConfig `mapstructure:"daemonTask"`
}

// scheduler daemon task config
type SchedulerModuleDaemonTaskConfig struct {
	Numerator map[string]time.Duration `mapstructure:"numerator"`
	Watcher   map[string]time.Duration `mapstructure:"watcher"`
}

// GetValue get value from config file
func GetValue[T any](key string, def T, getter ...func(string) T) T {
	if !slices.Contains(keys, strings.ToLower(key)) {
		return def
	}

	if len(getter) != 0 {
		return getter[0](key)
	}

	value := viper.Get(key)
	if value == nil {
		logger.Warnf("Null configuration item(%s) was found! Check whether it is correct", key)
		return def
	}

	if reflect.TypeOf(value).Kind() == reflect.Slice {
		valueSlice := reflect.ValueOf(value)

		// Create a new slice with the same type as the default value
		resultSlice := reflect.MakeSlice(reflect.TypeOf(def), valueSlice.Len(), valueSlice.Len())

		// Iterate through the slice and set the values
		for i := 0; i < valueSlice.Len(); i++ {
			elem := valueSlice.Index(i).Interface()

			// Check if the element type matches the default slice element type
			if reflect.TypeOf(elem).AssignableTo(reflect.TypeOf(def).Elem()) {
				resultSlice.Index(i).Set(reflect.ValueOf(elem))
			} else {
				panic(fmt.Sprintf("element of type %T is not assignable to type %T", elem, reflect.TypeOf(def).Elem()))
			}
		}

		return resultSlice.Interface().(T)
	}

	return value.(T)
}

// global config
var GlobalConfig *BMWConfig

// set the default value for different key
func setDefaultValue() {
	viper.SetDefault("service.mode", "release")
	// log basic config default value
	viper.SetDefault("log.level", "info")
	viper.SetDefault("log.path", "./bmw.log")
	viper.SetDefault("log.maxSize", 200)
	viper.SetDefault("log.maxAge", 1)
	viper.SetDefault("log.maxBackups", 5)
	// broker redis basic config default value
	viper.SetDefault("broker.redis.db", 0)
	viper.SetDefault("broker.redis.dialTimeout", 10*time.Second)
	viper.SetDefault("broker.redis.readTimeout", 10*time.Second)
	// store redis basic config default value
	viper.SetDefault("store.redis.db", 0)
	viper.SetDefault("store.redis.dialTimeout", 10*time.Second)
	viper.SetDefault("store.redis.readTimeout", 10*time.Second)
	viper.SetDefault("store.redis.keyPrefix", "bmw")
	// store dependent redis basic config default value
	viper.SetDefault("store.dependentRedis.db", 0)
	viper.SetDefault("store.dependentRedis.dialTimeout", 10*time.Second)
	viper.SetDefault("store.dependentRedis.readTimeout", 10*time.Second)
	// consul service basic config default value
	viper.SetDefault("store.consul.srvName", "bmw")
	viper.SetDefault("store.consul.port", 8500)
	viper.SetDefault("store.consul.tag", []string{"bmw"})
	// mysql basic config default value
	viper.SetDefault("store.mysql.port", 3306)
	viper.SetDefault("store.mysql.user", "root")
	viper.SetDefault("store.mysql.charset", "utf8")
	viper.SetDefault("store.mysql.maxIdleConnections", 10)
	viper.SetDefault("store.mysql.maxOpenConnections", 100)
	viper.SetDefault("store.mysql.debug", false)
	// task basic config
	viper.SetDefault("taskConfig.metadata.metricDimension.maxMetricsFetchStep", 500)
	// TODO: need replace by `timeSeriesMetricExpiredSeconds`
	viper.SetDefault("taskConfig.metadata.metricDimension.timeSeriesMetricExpiredDays", 30)
	// es basic config
	viper.SetDefault("store.es.esRetainInvalidAlias", false)
	// default queue
	viper.SetDefault("worker.queues", []string{"default"})
	viper.SetDefault("worker.concurrency", 0)
	viper.SetDefault("worker.healthCheck.interval", 3*time.Second)
	viper.SetDefault("worker.healthCheck.duration", 5*time.Second)
	viper.SetDefault("worker.daemonTask.maintainer.interval", 1*time.Second)
	viper.SetDefault("worker.daemonTask.maintainer.tolerateCount", 60)
	// scheduler daemon task config default value
	viper.SetDefault("scheduler.watcher.chanSize", 10)
	viper.SetDefault("scheduler.daemonTask.numerator.interval", 60*time.Second)
	viper.SetDefault("scheduler.daemonTask.watcher.workerWatchInterval", 1*time.Second)
	viper.SetDefault("scheduler.daemonTask.watcher.taskWatchInterval", 1*time.Second)
}

// InitConfig This method is used to refresh the configuration
// and should only be called once in the project.
// The purpose of this method is not private is that it can be called in the test file.
func InitConfig() {
	viper.SetConfigFile(FilePath)

	if err := viper.ReadInConfig(); err != nil {
		logger.Fatalf("read config file: %s error: %s", FilePath, err)
	}

	// set default value
	setDefaultValue()

	viper.AutomaticEnv()
	viper.SetEnvPrefix(EnvKeyPrefix)
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// map to struct config
	var config BMWConfig
	if err := viper.Unmarshal(&config); err != nil {
		logger.Fatalf("unmarshal config file: %s error: %s", FilePath, err)
	}
	GlobalConfig = &config
}
