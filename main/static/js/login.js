
new QWebChannel(qt.webChannelTransport, function(channel) {
    window.st = channel.objects.st;
  })

//LogInFun
function LogInFun(){
    window.st.onTestJsFunc(10, "ak");
    window.st.buildVersion = "12";
    console.log(window.st.buildVersion);
}