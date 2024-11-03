// page_left
window.onload = function () {
    var resize = document.getElementById("page_resize");
    var left = document.getElementById("page_left");
    var right = document.getElementById("page_right");
    var wrapper = document.getElementById("page_wrapper");
    resize.onmousedown = function (e) {
        var startX = e.clientX;
        resize.left = resize.offsetLeft;
        document.onmousemove = function (e) {
            var endX = e.clientX;

            var moveLen = resize.left + (endX - startX);
            var maxT = wrapper.clientWidth - resize.offsetWidth;
            if (moveLen < 150) moveLen = 150;
            if (moveLen > maxT - 150) moveLen = maxT - 150;

            resize.style.left = moveLen;
            left.style.width = moveLen + "px";
            right.style.width = (wrapper.clientWidth - moveLen - 2) + "px";
        }
        document.onmouseup = function (evt) {
            document.onmousemove = null;
            document.onmouseup = null;
            resize.releaseCapture && resize.releaseCapture();
        }
        resize.setCapture && resize.setCapture();
        return false;
    }
}
// pagr_left
//检索设备
// $("#queryDevice").on('shown.bs.select',function(e){
// 	$('#queryDevice').prev().find("input").keyup(function(){
// 		$('#queryDevice').prev().find("input").attr('id',"deviceInput"); //为input增加id属性
// 		console.log($('#deviceInput').val()); //获取输入框值输出到控制台
// 		var deviceInput = $('#deviceInput').val();
// 		var deviceStr="" ;
// 		for(var i=0; i<8; i++){
// 			deviceStr+="<option  data-icon='glyphicon glyphicon-heart' data-tokens='"+i+"'> 设备"+i+"</option>"; 
// 		}
// 		$("#queryDevice").html("");
// 		$('#queryDevice').append(deviceStr);
// 		$('#queryDevice').selectpicker('refresh');
// 	})
// });


