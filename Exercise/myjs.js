
function slideshow(){
    var banner=document.getElementById("banner");
    var ulis=banner.getElementsByTagName("li");
    //console.log(lis);
    var ol=document.getElementById("dot");
    var olis=ol.getElementsByTagName("li");
    //console.log(olis);
    for(var i=0;i<olis.length;i++){
        olis[i].index=i;
        olis[i].onmouseover=function(){
            for(j=0;j<olis.length;j++){
                olis[j].className="";
                ulis[j].className="";
                ulis[this.index].style.transition="all 2s ease 0.2s"
            }
            this.className="olActive";
            ulis[this.index].className="liActive";
            ulis[this.index].style.transition="all 2s ease 0.2s"
        }
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


addLoadEvent(slideshow);
addLoadEvent(tabNews);