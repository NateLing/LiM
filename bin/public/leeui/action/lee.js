(function(window,document){
    "use strict";
    var Lee = function(){};
    Lee.prototype={
        constructor:Lee,
        /**
         * 功能:向Lee注册插件
         * @param name
         * @param plugin
         */
        register:function(name,plugin){

        }
    };
    window.lee = new Lee();
})(window,window.document);