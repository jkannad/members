document.getElementById("country").addEventListener("change", function() {
    let val = document.getElementById("country").value;
    const stateDropDown = document.getElementById("state");
    var i, L = stateDropDown.options.length - 1;
    for(i = L; i >= 0; i--) {
        stateDropDown.remove(i);
    }
    fetch('/member/getstates/'+val)
        .then(response => response.json())
        .then(data => {
            let option = document.createElement("option");
            option.setAttribute("value", "");
            let optionText = document.createTextNode("Choose...");
            option.appendChild(optionText);
            stateDropDown.appendChild(option);  
            for(let i in data){
                let option = document.createElement("option");
                option.setAttribute("value", data[i].isoCode);
                let optionText = document.createTextNode(data[i].name);
                option.appendChild(optionText);
                stateDropDown.appendChild(option);
            }
        })
    let dialCode = document.getElementById("dial_code")
    if (dialCode) {
        fetch('/member/getdialcode/'+val)
        .then(response => response.json())
        .then(data => {
            console.log(data.dial_code)
            document.getElementById("dial_code").value = data.dial_code
            console.log("Dial Code text value:" + document.getElementById("dial_code").value)
            document.getElementById("dial_code_span").innerHTML = data.dial_code
        })
    }
})

document.getElementById("state").addEventListener("change", function() {
    let country = document.getElementById("country").value;
    let state = document.getElementById("state").value;
    const cityDropDown = document.getElementById("city");
    var i, L = cityDropDown.options.length - 1;
    for(i = L; i >= 0; i--) {
        cityDropDown.remove(i);
    }
    fetch('/member/getcities/'+country+'/'+state)
        .then(response => response.json())
        .then(data => {
            let option = document.createElement("option");
            option.setAttribute("value", "");
            let optionText = document.createTextNode("Choose...");
            option.appendChild(optionText);
            cityDropDown.appendChild(option);
            for(let i in data){
                let option = document.createElement("option");
                option.setAttribute("value", data[i].name);
                let optionText = document.createTextNode(data[i].name);
                option.appendChild(optionText);
                cityDropDown.appendChild(option);
            }
        })
})