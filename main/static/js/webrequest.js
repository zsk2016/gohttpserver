
document.write('<script type="text/javascript" src="/js/getotherdata.js"></script>')

function WebRequestGet(url){
  xhr = new XMLHttpRequest();
  xhr.open("GET", url, true);
  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4 && xhr.status == 200) {
      var responseText = xhr.responseText;
    }
  }
  xhr.send(null)
}

function WebRequest(webtype, type, normalUrl = "") {
  url = "/";
  url += type;
  if(type == "BufPay"){
    url = normalUrl;
  }
  xhr = new XMLHttpRequest();
  xhr.open("post", url, true);
  if(webtype == 'www'){
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
  }else if(webtype == 'json'){
    xhr.setRequestHeader("Content-Type", "application/json");
  }
  var data;
  if (type === "LogIn") {
    data = getLogInData();
  }else if(type == "SignIn"){
    data = getSignInData();
  }else if(type == "GetOrderByUser"){
    data = getOrderData();
  }else if(type == "GetRealize"){
    data = getRealize();
  }else if(type == "GetQRCodeUrl"){
    data = getQRCodeUrl();
  }else if(type == "BufPay"){
    data = getBufPayData();
  }
  xhr.onreadystatechange = function () {
    if (xhr.readyState == 4 && xhr.status == 200) {
      var jsonObj = JSON.parse(xhr.responseText);
      if (type === "LogIn"){
        setRetLogInData(jsonObj);
      }else if(type == "SignIn"){
        setRetSignInData(jsonObj);
      }else if(type == "GetOrderByUser"){
        setOrderData(jsonObj);
      }else if(type == "GetRealize"){
        setRealize(jsonObj);
      }else if(type == "GetQRCodeUrl"){
        setQRCodeUrl(jsonObj);
      }else if(type == "BufPay"){
        setBufPayData(jsonObj);
      }
    }
  }
  xhr.send(data);
}