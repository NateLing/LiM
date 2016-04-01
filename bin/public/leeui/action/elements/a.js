/**
 * Created by LingR on 2016/3/30.
 */
(function (document,lee) {
    var aElements = document.querySelectorAll('a');
    if (aElements instanceof NodeList) {
        var len = aElements.length,
            i = 0;
        for (; i < len; ++i) {
            (function (a) {
                var aHref = a.getAttribute('href');
                aHref && a.setAttribute('href', 'javascript:void(0);');
                a.addEventListener('click', function (evt) {});
            })(aElements[i]);
        }
    }
})(window.document,window.lee);

