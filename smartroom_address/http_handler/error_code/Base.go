package errCode

const (
	SUCCESS    = 200 //成功
	FAILURE    = 2   //失败
	SERVER_ERR = 500 //服务器错误

	//客户端错误
	BAD_REQUEST                  = 4000 // 数据格式错误
	LOGIN_USER_UN_EXTIST         = 4001 // 用户不存在
	LOGIN_PASSWORD_ACCOUNT_ERROR = 4002 // 密码或者账号错误
	LOGIN_NO_TOKEN               = 4003 // 未登录,请重新登陆
	LOGIN_NO_TOKEN_TIMEOUT       = 4005 //令牌已失效,请重新登陆
	DATA_EXIST                   = 4006 //数据已存在,重复插入
	JSON_MARSHAL_ERROR           = 4007 //JSON编码错误
	JSON_UNMARSHAL_ERROR         = 4008 //JSON解码错误
	HTTP_VERSION_ERROR           = 4009 //HTTP请求版本错误
	HTTP_POST_COMMAND_ERROR      = 4010 //HTTP请求命令错误
	USER_STATE_DISABLE           = 4011 //此用户不可用
	LOGIN_ON_OTHER_CLIENT        = 4012 //账号​在其它地点登录

	PAR_PARAMETER_NOT_LEGAL   = 4012 //参数不合法
	DATA_CONVERSION_EXCEPTION = 4013 //数据转换异常
	UNAUTHORIZED_OPERATION    = 4014 //无操作权限
	PAR_PARAMETER_IS_NULL     = 4015 //参数为空
	LOGIN_USER_NO_PERMISSION  = 4016 //用户没有登录权限
	ADD_ERR                   = 4017 //数据插入失败
	UPDATE_ERR                = 4018 //数据修改失败
	DELETE_ERR                = 4019 //数据删除失败
	PASSWORD_ERR              = 4020 //密码错误
	SMSCODE_ERR               = 4021 //验证码错误

	PAR_PARAMETER_TIME_ERROR = 4020 //时间格式不正确或者范围不正确   后来添加的

	ADD_ERROR    = 4040 //添加失败
	DATA_BE_USED = 4055 //数据已被使用
	DATA_TOO_BIG = 4060 //数据量过大
	//服务端错误
	SERVICE_BUSSY     = 5000 //服务器繁忙
	TIME_OUT          = 5001 //服务器查询超时
	DATA_NULL         = 5002 //数据不存在
	MGROBJ_FIND_ERROR = 5010 //获取设备信息失败
	GRPC_CONN_ERR     = 5011 //grpc连接失败
	GRPC_METHOD_ERR   = 5012 //GRPC方法调用失败

	UNKOWN_ERROR = -1 // "未知错误

	ROLE_CHECK_UESR_EXIST = 1601 //角色组下存在user

	//数据总览相关错误码
	PANDECT_GRPC_ERROR = 5101 //GRPC获取数据错误

	//获取机房概览数据错误码
	REAL_TIME_BASE_DATA_ERROR = 5201 //未获取到机房概览数据

	//报警事件错误码
	ALARM_EVENT_LIST_OGJ_NULL_ERROR = 5301 //查询同设备同监控项报警列表 设备表示同意为空
	ALARM_EVENT_LIST_NULL_ERROR     = 5302 //查询同设备同监控项报警列表 未获取到列表数据
	ALARM_CONFIRM_PARAM_ERROR       = 5303 //确认报警事件请求参数错误
	ALARM_CONFIRM_FAILED            = 5304 //确认报警事件失败
	ALARM_ASSIGN_PARAM_ERROR        = 5305 //指派报警事件请求参数错误
	ALARM_ASSIGN_FAILED             = 5306 //指派报警事件失败
	ALARM_LIST_PARAM_ERROR          = 5307 //获取报警事件列表参数错误

	//
	DATA_ERR = 502 //数据不存在
)
