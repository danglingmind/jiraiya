var uservalid = false;

function addRow() {
    var id = document.getElementById("inputs").childElementCount;
    var newRow = '<div id="new_input' + id + '">' +
        '<input type="text" name="name' + id + '" id="new_name' + id + '" placeholder="URL Name">' +
        '<input type="URL" name="url' + id + '" id="new_url' + id + '" placeholder="Enter URL">' +
        '</div>';

    document.getElementById("inputs").insertAdjacentHTML('beforeend', newRow);
}

function validateUser() {
    var letters = /^[0-9A-Za-z]+$/;
    userid = document.getElementById("userid").value;
    if (userid.length == 0) {
        warning("bad", "Comeon :/");
        uservalid = false;
    } else if (!userid.match(letters)) {
        warning("bad", "BAD UserID (only characters, numbers & _, . are allowed)");
        uservalid = false;
    } else if (userid.length < 5) {
        warning("bad", "Min Length is 5 characters");
        uservalid = false;
    } else {
        uservalid = true;
        warning("good", "Valid UserID");
    }
}

async function signInButton() {
    userid = document.getElementById("userid").value;
    if (uservalid) {
        var userpresesent = await userExists(userid);
        if (userpresesent) {
            passwordinput = '<input type="password" name="password" id="password" placeholder="Enter Password">';
            document.getElementById("useriddiv").insertAdjacentHTML("beforeend", passwordinput);

            document.getElementById("buttondiv").innerHTML =
                '<input type="button" id="signin" value="Sign In" onclick="signIn()">';
        } else {
            warning("bad", "Not Registered with us");
        }
    }

}

async function signUpButton() {
    userid = document.getElementById("userid").value;
    if (uservalid) {
        var userpresesent = await userExists(userid);
        if (userpresesent) {
            warning('bad', "Sign In, you are aleady registered");
        } else {
            // prompt for passwords
            passwordinput1 = '<br><input type="password" name="password1" id="password1" placeholder="Enter Password">';
            passwordinput2 = '<br><input type="password" name="password2" id="password2" placeholder="Re-Enter Password">';
            document.getElementById("useriddiv").insertAdjacentHTML("beforeend", passwordinput1);
            document.getElementById("useriddiv").insertAdjacentHTML("beforeend", passwordinput2);
            // change the button to just signup 
            document.getElementById("buttondiv").innerHTML =
                '<input type="button" id="signup" value="Sign Up" onclick="signUp()">';
        }
    }

}

async function signIn() {
    var userid = document.getElementById("userid").value;
    var password = document.getElementById("password").value;
    console.log(userid);
    console.log(password);
    if (password.length == 0) {
        warning("bad", "Enter Password");
        return
    }
    // check the login with API
    var loggedin = await login(userid, password);
    if (loggedin) {
        sessionStorage.setItem("userid", userid);
        window.location = "http://localhost:8000/render/cumul/user/inputurl/" + userid;
    } else {
        warning('bad', "Incorrect Password")
    }
}

async function login(userid, password) {
    var url = 'http://localhost:8000/cumul/' + userid + '/login/' + password;
    const response = await fetch(url);
    if (response.ok == true) {
        return true;
    } else {
        return false;
    }
}

async function signUp(userid) {
    userid = document.getElementById("userid").value;
    // check password
    pass1 = document.getElementById("password1").value;
    pass2 = document.getElementById("password2").value;
    if (pass1 != pass2) {
        warning("bad", "Passwords Did not match!!");
    } else {
        // create new user and redirect to next page 
        url = 'http://localhost:8000/cumul/' + userid + '/new/' + pass1;
        const response = await fetch(url);
        if (response.ok == true) {
            warning("good", "");
            document.getElementById("useriddiv").innerHTML = '<h2 style="color: aliceblue;>Hurreeyy !!</h2>';
            document.getElementById("useriddiv").style.color = 'alicblue';
            wait(1500);
            sessionStorage.setItem("userid", userid);
            window.location = "http://localhost:8000/render/cumul/user/inputurl/" + userid;
            return true;
        } else {
            warning("bad", "Couldnt connect to server!!");
            return false;
        }
    }
}

async function userExists(userid) {
    var url = 'http://localhost:8000/cumul/' + userid + '/check';
    const response = await fetch(url);
    if (response.ok == true) {
        return true;
    } else {
        return false;
    }
}

async function loadSavedURLs() {
    console.log('called');
    userid = sessionStorage.getItem("userid");
    url = 'http://localhost:8000/cumul/' + userid + '/urls';
    const response = await fetch(url);
    if (response.ok) {
        var jsonResponse = await response.json();
        var urls = jsonResponse.urls;
        if (urls != null) {
            // var alredyInputCount = document.getElementById("inputs").childElementCount;
            var i = 0;
            for (var u in urls) {
                i++;
                var savedInput = '<div id="input' + i + '">' +
                    '<input type="text" name="name' + i + '" id="name' + i + '" placeholder="URL Name">' +
                    '<input type="URL" name="url' + i + '" id="url' + i + '" placeholder="Enter URL">' +
                    '</div>';
                document.getElementById("inputs").insertAdjacentHTML("beforeend", savedInput);
                document.getElementById("name" + i).value = u;
                document.getElementById("url" + i).value = urls[u];
            }
        } else {
            // if none of the url is stored then create one default input section
            var newInput = '<div id="new_input0">' +
                '<input type="text" name="name0" id="new_name0" placeholder="URL Name">' +
                '<input type="URL" name="url0" id="new_url0" placeholder="Enter URL">' +
                '</div>';
            document.getElementById("inputs").innerHTML = newInput;

        }
    }
}

async function save() {
    userid = sessionStorage.getItem("userid");
    // store the newly added urls only
    jsonReq = {};
    var inputs = document.getElementById("inputs").children;
    for (var i = 0; i < inputs.length; i++) {
        if (inputs[i].id.indexOf("new_") == 0) {
            var data = inputs[i].children;
            jsonReq[data[0].value] = data[1].value;
        }
    }
    // call api to store the urls
    console.log(jsonReq);
    var url = 'http://localhost:8000/cumul/' + userid;
    var response = await fetch(url, {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(jsonReq)
    });

    if (response.ok) {
        alert("Url saved");
    } else {
        alert("try to save again");
    }
}


function wait(ms) {
    var start = new Date().getTime();
    var end = start;
    while (end < start + ms) {
        end = new Date().getTime();
    }
}

function warning(nature, msg) {
    document.getElementById("warning").innerHTML = msg;
    if (nature == 'bad') {
        document.getElementById("warning").style.color = "crimson";
    } else {
        document.getElementById("warning").style.color = "green";
    }
}

function checkUserSession() {
    var userid = sessionStorage.getItem("userid");
    if (userid) {
        // redirect to addurl page 
        // window.location = "http://localhost:8000/render/cumul/user/inputurl/" + userid;
    } else {
        // if not user session then redirect to login page 
        window.location = "http://localhost:8000/render/cumul/user/register/";
    }
}

