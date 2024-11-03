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