let LENGTH_MIN_EMAIL = 5;
let LENGTH_MIN_CONTRASEÑA = 2;
let LENGTH_MIN_NOMBRE= 1;
let MIN_IDDIV = 0;//numero muy bajo para no confrontar a max_idLi
let MAX_IDLI= 200000;//numero muy alto para que no de problemas 


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

        //reiniciamos el array de status de listas
        for(i=0;i<statusBotonesLista.length;i++){

            statusBotonesLista.pop;
        }
        var idList = MIN_IDDIV;

        listas.forEach( (list)=>{
   
            html += `<li  class="list-group-item list" onclick="onShowOptionsList('${idList}','div_${idList}')">${list.Nombre}
            <div class="botonLista" id="div_${idList}"> <button type="button" class="btn btn-primary"  onclick="onShowElements('${list.Id}')"> Mostrar elementos </button>
            <button type="button" class="btn btn-danger"  onclick="onEliminateList('${list.Id}')"> Eliminar Lista</button>
            </div></li>`;
            statusBotonesLista.push('false');
            idList = idList +1;
        } );

        html += '</ul>';

        htmlBoton = '<button id="boton_nueva_lista" type="button" class="btn btn-danger ms-4 " onclick="nuevaLista()">Añadir nueva lista</button>'


        //ponemos los botones relacionados a las listas
        $(buttons).html(htmlBoton);
        $(buttons).removeClass("botonNuevaLista").addClass("botonNuevaListaMostrar");
        //escondemos los de los elementos
        $(buttonsElementos).css("display","none");
        $(lista).html(html);


    } ).fail(pintaFallo);
}

function pintarListadoElementos(idLista,claseCss){

    var html = "";

    var queryArgs = "";
    queryArgs += "id=" + idLista;

    //debo conseguir el listado de nombres de elementos que pertenecen a la lista
    $.get(`/elementos-lista?${queryArgs}`).done( (elementos)=>{

        if(elementos.length != 0){
            html += '<ul class="list-group">';

            var idDiv = MIN_IDDIV;
            var idLi = MAX_IDLI;
            //reiniciamos el array de status de elementos
            for(i=0;i<statusBotonesElemento.length;i++){
    
                statusBotonesElemento.pop;
            }
    
            elementos.forEach( (elemento)=>{
                
                
                html += `<li id="li_${idLi}" class="list-group-item ${claseCss}" onclick="onShowOptionsElement('${idDiv}','div_${idDiv}','li_${idLi}','${claseCss}','${elemento.Nombre}','${ elemento.Id}')">${elemento.Nombre}
                <div class="botonElemento" id="div_${idDiv}"> <button type="button" class="btn btn-primary"  onclick="onMarcarElemento('li_${idLi}','div_${idDiv}')"> Marcar Elemento </button>
                <button type="button" class="btn btn-danger"  onclick="onEliminateElement('${elemento.Id}','div_${idDiv}')"> Eliminar Elemento</button>
                </div> </li>`;
                idDiv = idDiv + 1;
                idLi = idLi -1;
                statusBotonesElemento.push('false');
                
            } );
    
            html += '</ul>';
            pintaResultado("Mostrando elementos");
        } else{
            pintaResultado("No hay elementos en la lista");
        }

        
        $(lista).html(html);
        listaActual = idLista;

        var botones = "";
        botones += `<button id="boton_añadir_elemento" type="button" class="btn btn-primary" onclick="añadirElemento('${idLista}')"> Añadir Nuevo Elemento</button>`;
        botones += '<button id="boton_disparo" type="button" class="btn btn-warning" onclick="atras()" >Atras</button>';

        //escondemos los botones relacionados a las listas
        $(buttons).removeClass("botonNuevaListaMostrar").addClass("botonNuevaLista");
        $(buttonsElementos).css("display","block");
        $(buttonsElementos).html(botones);

    } ).fail(pintaFallo);

}

function onMarcarElemento(idLi,idDiv){

    $("#" + idDiv).removeClass( "botonElementoMostrar" ).addClass( "botonElemento" );
    $("#" + idLi).removeClass( "element" ).addClass( "elementoMarcado" );

}

function onEliminateList(listId){

    //query args
    queryArgs ="";
    queryArgs += 'elemento=list' ;
    queryArgs += "&idLista=" + listId;
    debug("Lista ID: " + listId);

    //eliminar elemento de la base de datos
    $.get(`/borrar?${queryArgs}`).done( (respuestaGet)=>{
        pintaResultado(respuestaGet);
        pintarListadoListas();
    } ).fail( (respuestaGet)=>{
        debug("he entrado por el fail");
        pintaFallo(respuestaGet);
        pintarListadoListas();
    } );
}

function hideFormulario(id){

    $("#" + id).css("display","none");
}
function showFormulario(id){
    $("#" + id).css("display","block");
}

function atras(){

    //debe irse desde los elementos de la lista, a mostrar el listado de listas
    pintarListadoListas();
    pintaResultado("Retorno a listas!");
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
            resolve(false);
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

function onShowOptionsList(indice,idList){
     //comprobamos el estado de idlist
     var estado = statusBotonesLista[indice];
     if( estado == "true" ){
         $("#" + idList).removeClass( "botonListaMostrar" ).addClass( "botonLista" );
         console.log("ha entrado");
         statusBotonesLista[indice] = 'false';
     } else{
         $("#" + idList).removeClass( "botonLista" ).addClass( "botonListMostrar" );
         statusBotonesLista[indice] = 'true';
     }
}

function onShowOptionsElement(indice,idDiv,idLi,clase,nombre,elementoId){
    //mostramos los botones a elegir
    debug("id div:" + idDiv );
    var idLi = idLi;
    var clase = clase;
    var nombre = nombre;
    var elementoId = elementoId;
    //comprobamos el estado de idDiv
    var estado = statusBotonesElemento[indice];
    if( estado == "true" ){
        $("#" + idDiv).removeClass( "botonElementoMostrar" ).addClass( "botonElemento" );
        console.log("ha entrado");
        statusBotonesElemento[indice] = 'false';
    } else{
        $("#" + idDiv).removeClass( "botonElemento" ).addClass( "botonElementoMostrar" );
        statusBotonesElemento[indice] = 'true';
    }
    

    
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
            pintarListadoListas();
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

function onSubmitNewElement(){

    var nombre = recogerFormularioString("nombreElemento",LENGTH_MIN_NOMBRE);
    if (nombre == ""){
        var resultado = "Error al meter nombre";
        pintaResultado(resultado);
    } else{
        var queryArgs = "";
        queryArgs += "nombre=" + nombre;
        queryArgs += "&idLista=" + listaActual;

        $.get(`/nuevo-elemento?${queryArgs}`).done( (respuestaGet) =>{
            pintaResultado(respuestaGet);
            hideFormulario("nuevoElemento");
            pintarListadoElementos(listaActual,"element");
        }).fail( (respuestaGet)=>{
            pintaFallo(respuestaGet);
            pintarListadoElementos(listaActual,"element");
        } );
    }
}

function onShowElements(listaId){

    pintarListadoElementos(listaId,"element");
}

function onEliminateElement(idElemento){

    //query args
    queryArgs ="";
    queryArgs += "elemento=" + idElemento;
    debug("ELEMENTO ID: " + idElemento)

    //eliminar elemento de la base de datos
    $.get(`/borrar?${queryArgs}`).done( (respuestaGet)=>{
        pintaResultado(respuestaGet);
        pintarListadoElementos(listaActual,"element");
    } ).fail( (respuestaGet)=>{
        pintaFallo(respuestaGet);
        pintarListadoElementos(listaActual,"element");
    } );
}

function añadirElemento(idLista){
    showFormulario("nuevoElemento");
    listaActual = idLista;
}

function eliminarElemento(idLista){

    pintarListadoElementos(idLista,"element");
    listaActual = idLista;

}




//******************** Empieza el script

var statusUsuario = "";
var listaActual = "";
//si los botones del elemento estan en display:block,true
var statusBotonesElemento = [];
var statusBotonesLista = [];


registro().then( (resolve) =>{

    debug("He llegado hasta el final");
    if (resolve == false){

    } else{
        pintarListadoListas();
    }
} );
