function debug(s) {
    console.log("Debug: " + s);
}

function pintaResultado(datos) {
    if (!datos) {
        datos = "OK";
    }
    var html = `<div class="alert alert-primary" >
    ${datos}
  </div>`;

    $( "#resultado" ).html( html );
} 

function pintaFallo(response) {

    var html = `<div class="alert alert-danger" >
    ${response.responseText}
  </div>`;
    $( "#resultado" ).html(html);
}  

//muestra los nombres de las listas disponibles que hay para consultar
function pintarListado(){

    var html = ""

    //debo conseguir el listado de nombres de listas de la base de datos de mi usuario
    $.get("/listas").done( (nombres)=>{
        html += '<ul class="list-group">';

        nombres.forEach( (nombre)=>{
   
            html += `<li class="list-group-item">${nombre}</li>`;
            
        } );

        html += '</ul>';

        $(list).html(html);

    } ).fail(pintaFallo);
}

function hideFormularioAutenticate(){

    $("#registros").css("display","none");
}
function showFormularioAutenticate(){
    $("#registros").css("display","block");
}

function autentificar(){

    $.get("/next-step").done( (status) =>{

        statusUsuario = status;
        if (statusUsuario == "autenticate"){
            showFormularioAutenticate();
        } else{
            pintaFallo(status);
        }

    } ).fail(pintaFallo);
}

function registro(){

    return new Promise( (resolve) => {

        $.get("/register").done( (respuestaGet)=>{
            pintaResultado(respuestaGet);
            resolve(true);
        } ).fail( (respuestaGet)=>{
            pintaFallo(respuestaGet);
            autentificar();
            resolve(true);
        })
        
    } );
}

//cuando se le da al boton submit, se ejecuta esta funcion
function recogerEmailFormularioRegistro() {
    var input = $("#email").val();
    var email = input.trim();

    if ( email.length < 5){
        //no compruebo la seguridad de la contraseña ni si el email es @
        return "";
    }
    //tengo ya los datos como coordenadas en un array
    return email;
}

function recogerContraseñaFormularioRegistro() {

    input = $("#contraseña").val();
    var contraseña = input.trim();

    if (contraseña.length < 2 ){
        //no compruebo la seguridad de la contraseña ni si el email es @
        return "";
    }
    //tengo ya los datos como coordenadas en un array
    return contraseña;
}

function onSubmitAutenticate(){

    var email = recogerEmailFormularioRegistro()
    var contraseña = recogerContraseñaFormularioRegistro()
    if (email == "" || contraseña == ""){
        var resultado = "Error al meter email y contraseña"
        pintaResultado(resultado)
    } else{
        var queryArgs = ""
        queryArgs += "email=" + email + "&" + "contraseña=" + contraseña

        $.get(`/autenticate?${queryArgs}`).done( (respuestaGet) =>{
            pintaResultado(respuestaGet);
            hideFormularioAutenticate();
        }).fail(pintaFallo);
    }
}




//******************** Empieza el script

var statusUsuario = "";


registro().then( () =>{

    debug("He llegado hasta el final");
    pintarListado();
} );
