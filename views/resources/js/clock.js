window.setInterval(ut, 1000);

function ut() {
  var d = new Date();
  document.getElementById("time").innerHTML = d.toLocaleTimeString();
  document.getElementById("date").innerHTML = d.toLocaleDateString();
  document.getElementById("time2").innerHTML = d.toLocaleTimeString();
  document.getElementById("date2").innerHTML = d.toLocaleDateString();
  document.getElementById("time3").innerHTML = d.toLocaleTimeString();
  document.getElementById("date3").innerHTML = d.toLocaleDateString();
}