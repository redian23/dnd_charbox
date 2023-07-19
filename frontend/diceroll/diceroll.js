var counter = 1;
function switchPlaces() {
    counter++
    if (counter %2){
        document.getElementById('main_place').style.display = "block";
    }else {
        document.getElementById('main_place').style.display = "none";
    }
}
function selectDice(){
    var s = document.getElementsByName('diceID')[0];
    s.addEventListener("change", selectDice);
    return s.options[s.selectedIndex].value
}
function diceCount() {
    var count = document.getElementById('number').value
    return count
}

var count = 1;
async function getDiceResult() {

    let dice_name = selectDice()
    let count = diceCount()

    const response = await fetch(`${window.location.protocol}//${window.location.hostname}:${window.location.port}/api/v1/roll?dice_name=${dice_name}&count=${count}`);
    const json = await response.json();
    let data = JSON.stringify(json);

    showElements()
    getRollSumStepsList(data);
    writeToHistory(data);
    writeResultInLabel(data)

}
function showElements() {
    document.getElementById("historyEmpty").innerHTML = "";
    document.getElementById("clearHistoryButton").style.display = "block";
}

function writeResultInLabel(data) {
    document.getElementById("diceRequest").innerHTML = diceCount() +"x"+ selectDice();
    document.getElementById("diceResult").innerHTML = JSON.parse(data)["total"];
}

function writeToHistory(data) {
    document.getElementById("historyList").innerHTML =count++ +". "+ diceCount()+"x"+ selectDice()
        + " Result: " + JSON.parse(data)["total"] + "<br>" + document.getElementById("historyList").innerHTML;
}

function getRollSumStepsList(data) {
    document.getElementById("rollSumHeader").style.display = "block";
    document.getElementById("rollSumTable").innerHTML = "";
    arr = JSON.parse(data)["roll_steps"];
    for(let i = 0; i < arr.length; i++){
        document.getElementById("rollSumTable").innerHTML += JSON.stringify(arr[i]["roll_number"]) +". --> "+ JSON.stringify(arr[i]["result"]) + "<br>" ;
    }
}
function plus(count){
    let value = parseInt(document.getElementById('number').value, 10);
    value = isNaN(value) ? 0 : value;
    value += count ;
    if (value >= 100) {
        value = 100
    }
    document.getElementById('number').value = value;
}
function minus(count){
    let value = parseInt(document.getElementById('number').value, 10);
    value = isNaN(value) ? 0 : value;
    value -= count ;
    if (value <= 1) {
        value = 1
    }
    document.getElementById('number').value = value;
}

function clearHistoryList() {
    document.getElementById("historyList").innerHTML = "";
}
function openTab(evt, tabName) {
    var i, x, tablinks;
    x = document.getElementsByClassName("content-tab");
    for (i = 0; i < x.length; i++) {
        x[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tab");
    for (i = 0; i < x.length; i++) {
        tablinks[i].className = tablinks[i].className.replace("is-active", "");
    }
    document.getElementById(tabName).style.display = "block";
    evt.currentTarget.className += " is-active";
}