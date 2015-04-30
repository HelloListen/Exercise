
function slideshow(){
    var ol=document.getElementById("dot");
    var olis=ol.getElementsByTagName("li");
    var div=document.getElementById("banner");
    var ulis=div.getElementsByTagName("li");
    var ul=document.getElementsByClassName("banner-item")[0];
    ul.innerHTML=ul.innerHTML+ul.innerHTML;
    //console.log(ul);
    var left=parseInt(ul.style.left)||0;
   // console.log(left);
    var timer=setInterval(function(){
        ul.style.left=left+"px";
        //olis[i].className="olActive";
       // ul.style.transition="all 1s ease 0.2s";
            if(left<-2240){
                left=0;
            }
        left=left-560;
    },2000);
}


addLoadEvent(slideshow);