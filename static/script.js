const list = $("#list");
//console.log(list);
const form = $("#form");
//console.log(form);

//Function nn keypress
form.keypress(function(event){

    //If not 13
    if(event.keyCode != 13){
        return;
    }

    event.preventDefault();

    //Hold onto it before it gets wiped
    const text = form.val();
    form.val(""); 

    //If text is invalid, return
    if(!text.trim()){
        return;
    }

    list.append("<li id = 'input'  align='right' class=\"list-group\">" + "You: " + text + "</li class=\"list-group\">");

    //Code adapted from: http://api.jquery.com/jquery.ajax/
    //jQuery get request.
    $.get("/ask", {input:text})
         //Sends parameters
        .done(function(elizaPattern){ 
            //Executes the response from server
            list.append("<li id='output' align='left' class=\"list-group\">" + "Eliza: "+ elizaPattern + "</li class=\"list-group\">");
        }).fail(function(){ 
            // fail runs if there was any error.
            list.append("<li>Sorry!</li>");
        });
        window.scrollTo(0, document.body.scrollHeight);
});