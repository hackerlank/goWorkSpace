package config

const (
	CONF_FILE_PATH = "core/config.json"
	ZK_CRAWLER_ROOT_PATH = "/dspider"
	ZK_CRAWLER_MASTER_PATH = "/dspider/master"
	ZK_CRAWLER_MASTER_HEARTBEAT_PATH = "/dspider/master/heartbeat"
	ZK_CRAWLER_SCHEDULER_PATH = "/dspider/scheduler"
	ZK_CRAWLER_SCHEDULER_HEARTBEAT_PATH = "/dspider/scheduler/heartbeat"
	ZK_CRAWLER_DP_PATH = "/dspider/dp"
	// actual node path: "/dspider/dp/{id}"
	// actual node heartbeat path: "/dspider/dp/{id}/heartbeat"
	//ZK_CRAWLER_DP_HEARTBEAT_PATH = "/dspider/dp/heartbeat"
	ZK_CRAWLER_SINK_PATH = "/dspider/sink"
	ZK_CRAWLER_SINK_HEARTBEAT_PATH = "/dspider/sink/heartbeat"

	ZK_CRAWLER_MASTER_HEARTBEAT = 60
	ZK_CRAWLER_SCHEDULER_HEARTBEAT = 60
	ZK_CRAWLER_DP_HEARTBEAT = 60
	ZK_CRAWLER_SINK_HEARTBEAT = 60

	ZK_CRAWLER_HEARTBEAT_KEY = "heartbeat_ts"
	ZK_CRAWLER_HOSTNAME_KEY = "hostname"
	ZK_CRAWLER_IP_KEY = "ipv4"
	ZK_CRAWLER_START_KEY = "start_ts"
)
