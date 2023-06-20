<h1>Get Temperature By City</h1>

<!-- <form action="http://localhost:3000/temperature/Mississauga" method="POST">
<div><input type="submit" value="Save"></div>
</form>-->

<form action="/results" id="cityform">
    <div><h4>Enter a city name</h4></div>
    <div><input type="text" class="form-control" id="userinput" placeholder="City Name"></input></div>
    <button type="submit">Submit</button>
</form>

<script>
    let form = document.getElementById("cityform");

    form.addEventListener("submit", (e) => {
        e.preventDefault();

        let city = document.getElementById("userinput");

        if (city.value == "") {
            alert("Please enter a city name");
        } else {
            // perform operation with form input
            //alert("This form has been successfully submitted!");
            //console.log(`the city entered is ${city.value}`);
            

            
            //city.value = "";
        }
    });
</script>