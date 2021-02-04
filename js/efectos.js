 
 $(document).ready(function(){
$("#logos").hide();
$("#videos").hide();
   
});
$("#mostrarL").click(function(){
    $("#logos").show(1000);
    $("#mostrarL").hide();
       
});

$("#mostrarV").click(function(){
    $("#videos").show(1000);
    $("#mostrarV").hide();
       
});
