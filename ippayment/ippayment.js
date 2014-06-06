   //$(document).ready(function(){
    jQuery.noConflict();
   	jQuery(document).ready(function($){	           
     // Removed the unnecessary `$(function() { ...` here and the matching closing bits at the end
     if($("#block").text() == 'Yes') {
     
     	$("#sb").hide();
     	$("#block").css("color","red");
     	
     
     }
          
     
     $("#sb").click(function(e){
        e.preventDefault();
        
        var tel =  $("#tel").text()
        //tel = encodeURIComponent(tel);//url encodes data
        $.ajax({
           type: "GET",
           url: "/blocktel",
           data: {'tel': tel},
           //dataType: "json",
           success: function(data) {
           	  alert('Block '+tel +" Successful!");	
              $("#sb").hide();
              $("#block").text("Yes")
              $("#block").css("color","red");
             alert('Block '+tel);
       }
       });
    });
    });
