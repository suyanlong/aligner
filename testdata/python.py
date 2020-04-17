from gos_app import ts


# 节点类型
class nodeType:
    SN = "2"  # 超级节点;如果backUp为true为超级节点,false为备用超级节点
    IN = "1"  # 产业节点
    UN = "0"  # 用户节点


# 用户类型
class userType:
    person = "0"  # 个人
    company = "1"  # 企业


# 明细类型
class GosTxType:
    gosMallPay = "-10"  # 公链商城支出
    appPay = "-9"  # 应用支出
    ecPay = "-8"  # 商城支出
    addPool = "-7"  # 补充资源池
    cityFee = "-6"  # 城市中心奖励支出
    taskFee = "-5"  # 任务费
    vote = "-4"  # 投票
    votePay = "-3"  # 发起人支出,投票奖励发放
    pay = "-2"  # 支出
    contribute = "-1"  # 贡献
    invite = "1"  # 邀请奖励
    task = "2"  # 任务奖励
    ec = "3"  # 商城奖励
    first = "4"  # 首次赠送
    activity = "5"  # 做活动奖励
    poolEarnings = "6"  # 算例资源收益
    profit = "7"  # 分红收益
    income = "8"  # 收入
    sponsorIncome = "9"  # 发起人收入,投票-资源池奖励
    voterIncome = "10"  # 投票人收入,投票奖励收入
    voteExpire = "11"  # 投票到期
    wholesale = "12"  # 批发
    appIncome = "13"  # 应用收入
    taskPlatform = "14"  # 平台任务奖励
    payBack = "15"  # 退回


# 状态
class State:
    doing = "0"  # 未完成
    success = "1"  # 成功
    failure = "2"  # 失败


# 错误
class Error:
    correct = 0  # 正确
    fault = 0  # 错误


# 活动事件
class eventType:
    city = "1"  # 城市中心活动奖励


# 事项
Matter = {
    "0": "给公链算力中心提供算力资源支持"
}


# 任务类型
class TaskType:
    main = "0"  # 主线任务
    ec = "1"  # 商城任务
    train = "2"  # 培训会任务
    re = "3"  # 推荐任务


# 消息类型
class MessageType:
    system = "0"  # 系统
    user = "1"  # 用户


# 支付数量
class PayAmount:
    trainAmount = ts.get("trainAmount")  # 培训费
    if trainAmount:
        train = float(trainAmount)
    else:
        train = 9800

    trainAmount2 = ts.get("trainAmount2")  # 培训费
    if trainAmount2:
        train2 = float(trainAmount2)
    else:
        train2 = 4800

    trainGos = ts.get("trainGos")  # 培训费
    if trainGos:
        gos = float(trainGos)
    else:
        gos = 90

    trainGos2 = ts.get("trainGos2")  # 培训费
    if trainGos2:
        gos2 = float(trainGos2)
    else:
        gos2 = 100


class PayType:
    pay = "-2"  # 支出


class UpgradeType:
    UTI = 1  # UN >> IN
    UTS = 2  # UN >> SN
    ITS = 3  # N >> SN


# 1. 4月3日 2.8月15日
SFB = int(ts.get("SFB")) if ts.get("SFB") else 8100  # 超级节点首次赠送
IFB = int(ts.get("IFB")) if ts.get("IFB") else 1620  # 产业节点首次赠送

# aligner disable
class Bonus:
    FUT_IB = IFB  # 用户节点升产业节点首次赠送
    FUT_SB = SFB  # 用户节点升超级节点首次赠送
    FIT_SB = SFB - IFB  # 产业节点升超级节点首次赠送
    RIB_S = SFB * 0.2  # 根节点邀请奖励,SN
    RIB_I = IFB * 0.1  # 根节点邀请奖励,IN
    UTS = SFB * 0.1  # 用户节点升超级节点总邀请奖励
    UTI = IFB * 0.2  # 用户节点升产业节点总邀请奖励
    ITS = UTS - UTI  # 产业节点升超级节点总邀请奖励
    OIB_S = SFB * 1.3 / 9  # 运营邀请奖励,SN
    OIB_I = IFB * 1.3 / 9  # 运营邀请奖励,IN
    SIS = SFB * 0.1  # 超级节点邀请超级节点邀请奖励
    SII = IFB * 0.2  # 超级节点邀请产业节点邀请奖励
    III = IFB * 0.1  # 产业节点邀请产业节点奖励


class ApiType:
    CityCenter = "city_center"
    UpgradePackage = "upgrade_package"
    Rebate = "rebate"
    IterationApp = "iteration_app"


class ApiOperation:
    Notification = "notification"
    AcquireData = "acquire_data"

# aligner enable
class RebateCode:
    UIU_A = 10001000  # 用户节点推荐用户节点激活
    IIU_A = 1000100  # 产业节点推荐用户节点激活
    SIU_A = 10001002  # 超级节点推荐用户节点激活
    III_S = 10004100  # 产业节点推荐产业节点开户
    SII_S = 10004101  # 超级节点推荐产业节点开户
    SIS_S = 10004200  # 超级节点推荐超级节点开户
    RIS_S = 10004201  # 根节点推荐超级节点开户
    # 升级
    UIU_UI = 10002001  # 用户节点推荐用户节点升级产业节点
    IIU_UI = 10002002  # 产业节点推荐用户节点升级产业节点
    SIU_UI = 10002003  # 超级节点推荐用户节点升级产业节点
    UIU_US = 10002100  # 用户节点推荐用户节点升级超级节点
    IIU_US = 10002101  # 产业节点推荐用户节点升级超级节点
    SIU_US = 10002102  # 超级节点推荐用户节点升级超级节点
    III_US = 10002200  # 产业节点推荐产业节点升级超级节点
    SII_US = 10002201  # 超级节点推荐产业节点升级超级节点
    UII_US = 10002202  # 用户节点推荐产业节点升级超级节点
    # 算力资源池
    SN_PB = 10003000  # 超级节点算力资源池收益
    BSN_PB = 10003001  # 备用超级节点算力资源池收益
    # 投票（暂无）
    # 10004000  # 超级节点投票
    # 10004001  # 备用超级节点投票
