/**
 * Created by LingR on 2016/3/30.
 */
(function (d) {
    var aElements = d.querySelectorAll('a');
    if (aElements instanceof NodeList) {
        var len = aElements.length,i=0;
        for(;i<len;++i){
            (function(a){
                var aHref = a.getAttribute('href');
                aHref && a.setAttribute('href','javascript:return false;');
                a.addEventListener('click',function(){

                });
            })(aElements[i]);
        }
    }
})(window.document);