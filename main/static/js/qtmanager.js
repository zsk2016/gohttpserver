new QWebChannel(qt.webChannelTransport, function (channel) {
    window.bridge = channel.objects.bridge;
})
function onShowMsgBox() {
    if (bridge) {
        bridge.showMsgBox()
    }
}

function onLogInState(isOk){
    if (bridge) {
        bridge.onSetLogInState(isOk)
    }
}

function onLogUserInfo(userName, userId){
	if (bridge) {
		bridge.onSetUserInfo(userName, userId)
	}
}

function onSetCpuId(id){
	window.sessionStorage.setItem("CpuId", id);
}

function showAlert() {
    alert('this is web alert');
}