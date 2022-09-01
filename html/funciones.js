let LENGTH_MIN_EMAIL = 5;
let LENGTH_MIN_CONTRASEÑA = 2;
let LENGTH_MIN_NOMBRE= 1;

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
function pintarListadoListas(){

    var html = ""

    //debo conseguir el listado de nombres de listas de la base de datos de mi usuario
    $.get("/listas").done( (listas)=>{
        html += '<ul class="list-group">';


        listas.forEach( (list)=>{
   
            html += `<li class="list-group-item item" onclick="onShowElements(${list.Nombre},${list.Id})">${list.Nombre}</li>`;
        } );

        html += '</ul>';

        $(lista).html(html);

    } ).fail(pintaFallo);
}

function pintarListadoElementos(nombreLista,idLista){

    var html = "";

    var queryArgs = "";
    queryArgs += "nombre=" + nombreLista;
    queryArgs += "&id=" + idLista;

    //debo conseguir el listado de nombres de elementos que pertenecen a la lista
    $.get(`/elementos-lista?${queryArgs}`).done( (elementos)=>{
        html += '<ul class="list-group">';

        elementos.forEach( (elemento)=>{
   
            html += `<li class="list-group-item">${elemento.Nombre}</li>`;
            
        } );

        html += '</ul>';

        $(lista).html(html);

    } ).fail(pintaFallo);

}

function hideFormulario(id){

    $("#" + id).css("display","none");
}
function showFormulario(id){
    $("#" + id).css("display","block");
}

function autentificar(){

    $.get("/next-step").done( (status) =>{

        statusUsuario = status;
        if (statusUsuario == "autenticate"){
            showFormulario("registros");
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

//le pasas por argumentos el id del input y la longitud para comprobar
//siempre que sean strings
function recogerFormularioString(id,length) {

    input = $("#" + id).val();
    var inputTrim = input.trim();

    if (inputTrim.length < length ){
        return "";
    }
    //tengo ya los datos como coordenadas en un array
    return inputTrim;
}


function onSubmitAutenticate(){

    var email = recogerFormularioString("email",LENGTH_MIN_EMAIL);
    var contraseña = recogerFormularioString("contraseña",LENGTH_MIN_EMAIL);
    if (email == "" || contraseña == ""){
        var resultado = "Error al meter email y contraseña";
        pintaResultado(resultado);
    } else{
        var queryArgs = "";
        queryArgs += "email=" + email + "&" + "contraseña=" + contraseña;

        $.get(`/autenticate?${queryArgs}`).done( (respuestaGet) =>{
            pintaResultado(respuestaGet);
            hideFormulario("registros");
        }).fail(pintaFallo);
    }
}

function nuevaLista(){

    showFormulario("nuevaLista");
}

function onSubmitNewList(){

    var nombre = recogerFormularioString("nombre",LENGTH_MIN_NOMBRE);
    if (nombre == ""){
        var resultado = "Error al meter nombre";
        pintaResultado(resultado);
    } else{
        var queryArgs = ""
        queryArgs += "nombre=" + nombre

        $.get(`/crear-lista?${queryArgs}`).done( (respuestaGet) =>{
            pintaResultado(respuestaGet);
            hideFormulario("nuevaLista");
            pintarListadoListas();
        }).fail( (respuestaGet)=>{
            pintaFallo(respuestaGet);
            pintarListadoListas();
        } );
    }
}

function onShowElements(nombre,listaId){

    debug ("hola que tal");
    pintarListadoElementos(nombre,listaId);
}




//******************** Empieza el script

var statusUsuario = "";


registro().then( () =>{

    debug("He llegado hasta el final");
    pintarListadoListas();
} );
