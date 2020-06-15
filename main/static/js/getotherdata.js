//登录信息
function getLogInData() {
    var data = "LogIn=";
    dataInfo = {
        "UserName": document.getElementById("loginusername").value,
        "Passwd": document.getElementById("loginpassword").value,
        "CpuId": ""
    }
    data += JSON.stringify(dataInfo);
    return data;
}

function setRetLogInData(data) {
    if (data.Ok) {
        window.sessionStorage.setItem("UserName", data.value.UserName);
        window.sessionStorage.setItem("UserId", data.value.UserId);
        document.getElementById("revalue").innerHTML = "登录成功!";
        WebRequest('www', 'GetOrderByUser');
        //onLogInState(true);
    }
    else {
        document.getElementById("revalue").innerHTML = "登录失败!";
    }
}

//注册信息
function getSignInData() {
    var data = "SignIn=";
    dataInfo = {
        "UserName": document.getElementById("signinusername").value,
        "EmailAddr": document.getElementById("signinemail").value,
        "Passwd": document.getElementById("signinpassword").value,
        "Company": document.getElementById("signincompany").value,
        "Phone": document.getElementById("signinphone").value,
        "UserType": "1",
        "PcNum": "2",
        "CpuId": ""
    }
    data += JSON.stringify(dataInfo);
    return data;
}

function setRetSignInData(data) {
    if (data.Ok) {
        document.getElementById("signinrevalue").innerHTML = "注册成功!";
        window.location = "../../views/pages/index.html";
    }
    else {
        document.getElementById("signinrevalue").innerHTML = data.data;
    }
}

//获取订单
function getOrderData() {
    var data = "SignIn=";
    dataInfo = {
        "UserName": window.UserId,
        "UserId": window.UserId
    }
    data += JSON.stringify(dataInfo);
    return data;
}

function setOrderData(data) {
    if (data.Ok) {
        rTime = data.value.RemainTime;
        ifValid = data.value.IfValid;
        mustUpdate = data.value.MustUpdate;
        info = data.value.Info;
        if (mustUpdate == true) {
            ;
        } else if (rTime >= 0 && ifValid == true) {
            document.getElementById("revalue").innerHTML = "剩余" + rTime + "小时";
        } else if (ifValid == false) {
            document.getElementById("revalue").innerHTML = "本次免费登录!"
            setTimeout('onLogInState(true)', 2000);
        } else {
            document.getElementById("revalue").innerHTML = "余额不足, 请缴费!";
        }
    }
}

//获取业务功能
function getRealize() {
    var data = "SignIn=";
    dataInfo = "";
    data += JSON.stringify(dataInfo);
    return data;
}

function setRealize(data) {
    if (data.Ok) {
        arraySize = data.value.length;
        for (let i = 0; i < arraySize; i++) {
            console.log(data.value[i].Id);
        }
        window.sessionStorage.setItem("realizeValueKey", JSON.stringify(data.value));
        window.location.href = "/pages/pricelist.html";
    }
}

//获取购买链接
function getQRCodeUrl() {
    var data = "GetQRCodeUrl=";
    dataInfo = {
        "RealizeId": window.sessionStorage.getItem("currentRealize"),
        "UserName": window.sessionStorage.getItem("UserName"),
        "ReturnLinkaddr": "localhost:8000/",
    }
    data += JSON.stringify(dataInfo);
    return data;
}

function setQRCodeUrl(data) {
    if (data.Ok) {
        rdata = data.value;
        document.getElementById("qrcodeimg").src = rdata.QrcCode;
        var qrctype = ""
        if(rdata.QrcType == "wechat")
            qrctype = "微信支付";
        qrcInfo = "支付类型: "+qrctype + "\n订单号: "+rdata.OrderId + "\n价格: "+rdata.Price + "元";
        document.getElementById("qrcodeInfo").innerHTML = qrcInfo;
    }
}

//bufpay
function getBufPayData(){
    var data = "";
    bufData = {
        "name": window.bufinfo.Name,
        "pay_type": window.bufinfo.Pay_type,
        "price": window.bufinfo.Price,
        "order_id": window.bufinfo.Order_id,
        "order_uid": window.bufinfo.Order_uid,
        "notify_url": window.bufinfo.Notify_url,
        "return_url": window.bufinfo.Return_url,
        "sign": window.bufinfo.Sign,
    }
    data += JSON.stringify(bufData);
    return data;
}

function setBufPayData(data){
    ;
}