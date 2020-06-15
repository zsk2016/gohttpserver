
window.onload = function () {
    realizeValue = JSON.parse(window.sessionStorage.getItem("realizeValueKey"));
    var num = realizeValue.length;
    for (let i = 0; i < num; i++) {
        document.getElementById("feature" + (i + 1).toString()).innerHTML = realizeValue[i].Name;
        document.getElementById("price" + (i + 1).toString()).innerHTML = realizeValue[i].Value;
        document.getElementById("time" + (i + 1).toString()).innerHTML = realizeValue[i].Ptime;
    }
    document.getElementsByName('rd')[0].checked = true;
} 

function acceptBuy(){
    realizeValue = JSON.parse(window.sessionStorage.getItem("realizeValueKey"));
    var rd = document.getElementsByName('rd');
    for(let i=0; i<realizeValue.length; i++ )
    {
        if(rd[i].checked)
        {
            window.sessionStorage.setItem("currentRealize", realizeValue[i].Id);
            console.log(realizeValue[i].Id);
            break;
        }
    }
    WebRequest("www", "GetQRCodeUrl");
}