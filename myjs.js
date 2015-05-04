
function slideshow(){
    var banner=document.getElementById("banner");
    var ulis=banner.getElementsByTagName("li");
    //console.log(lis);
    var ol=document.getElementById("dot");
    var olis=ol.getElementsByTagName("li");
    //console.log(olis);
    autoPlay(0);
    for(var i=0;i<olis.length;i++){
        olis[i].index=i;
        olis[i].onmouseover=function(){
            clearInterval(timer);
            for(j=0;j<olis.length;j++){
                olis[j].className="";
                ulis[j].className="";
            }
            this.className="olActive";
            ulis[this.index].className="liActive";
        }
        olis[i].onmouseout=function(){
            autoPlay(this.index);
        }
    }
    function autoPlay(_index){
         timer=setInterval(function(){
            for(var k=0;k<olis.length;k++){
                olis[k].className="";
                ulis[k].className="";
            }
            ulis[_index].className="liActive";
            olis[_index].className="olActive";
             _index++;
            if(_index==olis.length){
                _index=0;
            }
        },2000);
    }
}


function tabNews(){
    var titles=document.getElementsByClassName("tab-title");
    //console.log(titles);
    var contents=document.getElementsByClassName("row-display-none");
    //console.log(contents);
    for(var i=0;i<titles.length;i++){
        titles[i].index=i;
        titles[i].onmouseover=function(){
            for(j=0;j<titles.length;j++){
                titles[j].className="col-md-4 tab-title";
                contents[j].className="row row-display-none";
            }
            this.className="spotlight col-md-4 tab-title";
            //this.style.transition="all 2s ease 0.2s";
            //console.log(this.index);
            //titles[this.index].className="spotlight col-md-4 tab-title";
            contents[this.index].className="row-active row row-display-none";
        }
    }
}

function blockQuote(){
    var wrap=document.getElementById("wrap");
    var section2=wrap.getElementsByClassName("section2")[0];
    //console.log(section2);
    var wrapcontent=wrap.getElementsByClassName("wrap-content")[0];
    console.log(wrapcontent);
    //wrapcontent.innerHTML=wrapcontent.innerHTML+wrapcontent.innerHTML;
    var left=parseInt(wrapcontent.style.left)||0;
    //console.log(left);
    var timer=setInterval(function(){
        wrapcontent.style.left=left+"px";
        left-=1100;
        if(left<-1100){
            left=0;
        }
        wrapcontent.style.transition="all 2s ease 0.2s";
    },5000);
}





addLoadEvent(slideshow);
addLoadEvent(tabNews);
addLoadEvent(blockQuote);