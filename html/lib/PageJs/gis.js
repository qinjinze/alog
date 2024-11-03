var gisGoogle = {
  marker: null,
  markers: [],
  path: null,
  infoWindow: null,
  initMap(divId) {
    map = new google.maps.Map(document.getElementById(divId), {
      center: { lat: 22.545573, lng: 114.111295 },
      zoom: 10
    });
  },
  addMarker(options) {
    if (options.id) {
      this.clearMarker(options.id);
    }
    var uluru = { lat: options.y, lng: options.x };
    var marker = new google.maps.Marker({
      position: uluru,
      map: map,
      icon: options.icon
    });
    marker.cid = options.id;
    this.markers.push(marker);
    map.setCenter(uluru);
    if (this.infoWindow) {
      this.infoWindow.close();
      this.infoWindow = null;
    }
    if (options.popup) {
      marker.addListener('click', function (e) {
        if (this.infoWindow) {
          this.infoWindow.close();
          this.infoWindow = null;
        }
        this.infoWindow = new google.maps.InfoWindow({
          content: options.popup()
        });
        this.infoWindow.open(map, marker);
      }.bind(this));
    }
  },
  clearMarker(id) {
    if (this.marker) {
      this.marker.setMap(null);
      this.marker = null;
    }
    if (this.infoWindow) {
      this.infoWindow.close();
      this.infoWindow = null;
    }
    if (this.markers.length > 0) {
      for (let i = this.markers.length - 1; i >= 0; i--) {
        var marker = this.markers[i];
        if (marker.cid == id) {
          marker.setMap(null);
          marker = null;
          this.markers.splice(i, 1);
          return i;
          break;
        }
      }
    }
  },
  addLine(options) {                                                                                                                                            
    var points = []; 
    var xmin = options.coords[0][0],
      ymin = options.coords[0][1],
      xmax = options.coords[0][0],
      ymax = options.coords[0][1];
    options.coords.forEach(data => {
      points.push({ lng: data[0], lat: data[1] });
      if (xmin > data[0]) xmin = data[0];
      if (ymin > data[1]) ymin = data[1];
      if (xmax < data[0]) xmax = data[0];
      if (ymax < data[1]) ymax = data[1];
    })
    var sy = {
      path: google.maps.SymbolPath.FORWARD_OPEN_ARROW,
      scale:1.5, //图标缩放大小
      strokeColor: '#fff', //设置矢量图标的线填充颜色
      strokeWeight: '1', //设置线宽
      fillColor: '#fff',
      fillOpacity: 0.8,
      strokeOpacity:1
    }
    this.path = new google.maps.Polyline({
      path: points,
      geodesic: true,
      strokeColor: '#0000FF',
      strokeOpacity: 1.0,
      strokeWeight: 6,
      icons: [{
        icon: sy,
        offset: '50px',
        repeat: '50px',
      }]
    });
    this.path.setMap(map);

    var bounds = new google.maps.LatLngBounds({ lng: xmin, lat: ymin }, { lng: xmax, lat: ymax });
    map.fitBounds(bounds);
  },
  clearLine() {
    if (this.path) {
      this.path.setMap(null);
      this.path = null;
    }

  },
  getLocation(options) {
    var geocoder = new google.maps.Geocoder;
    var pt = new BMap.Point(options.x, options.y);
    var latlng = { lng: options.x, lat: options.y }
    geocoder.geocode({ 'location': latlng }, function (results, status) {
      var address;
      if (results) {
        if (results[0])
          address = results[0].formatted_address
      }
      setTimeout(() => {
        options.callback(address);
      }, 300);

    });
  }
}

var gisBaidu = {
  marker: null,
  markers: [],
  path: null,
  infoWindow: null,
  initMap(divId) {
    //百度地图功能
    map = new BMap.Map(divId);    // 创建Map实例
    map.centerAndZoom(new BMap.Point(114.111295, 22.545573), 10);  // 初始化地图,设置中心点坐标和地图级别
    //添加地图类型控件
    map.addControl(new BMap.MapTypeControl({
      mapTypes: [
        BMAP_NORMAL_MAP,
        BMAP_HYBRID_MAP
      ]
    }));
    // 添加带有定位的导航控件
    var navigationControl = new BMap.NavigationControl({
      // 靠左上角位置
      anchor: BMAP_ANCHOR_TOP_LEFT,
      // LARGE类型
      type: BMAP_NAVIGATION_CONTROL_LARGE,
      // 启用显示定位
      enableGeolocation: true
    });
    // function showInfo(e){
    // 	alert(e.point.lng + ", " + e.point.lat);
    // }
    // map.addEventListener("click", showInfo);
    map.addControl(navigationControl);
    map.setCurrentCity("深圳");
    map.enableScrollWheelZoom(true);     //开启鼠标滚轮缩放

    //滑动事件
    map.addEventListener("mousemove", function (e) {
      var p = e.point;
      $("#coordTip").html(p.lng + ',' + p.lat);
    });
  },
  addMarker(options) {
    if (options.id) {
      this.clearMarker(options.id);
    }
    var point = new BMap.Point(options.x, options.y);
    // var marker = new BMap.Marker(point);
    var marker;
    if(options.icon){
      var iconObj = new BMap.Icon(options.icon, new BMap.Size(19, 31));
      marker = new  BMap.Marker(point, {
        icon:iconObj,
        offset: new BMap.Size(0, -15)
      });
    }else {
      marker = new BMap.Marker(point);
    }
    marker.cid = options.id;
    this.markers.push(marker);
    map.addOverlay(marker);
    map.centerAndZoom(point, 15);
    // if(options.zoom){
    //   map.centerAndZoom(point, 15);
    // }
    var opts = {
      width: 300,     // 信息窗口宽度
      height: 120,     // 信息窗口高度
      title: "",// 信息窗口标题
      enableMessage: true,//设置允许信息窗发送短息
      message: "1111",
    };
    if (this.infoWindow) {
      map.closeInfoWindow(this.infoWindow);
    }
    marker.addEventListener("click", function (e) {
      var content = options.popup(options.id);
      this.infoWindow = new BMap.InfoWindow(content, opts);  // 创建信息窗口对象 
      map.openInfoWindow(this.infoWindow, point); //开启信息窗口
    }.bind(this))
  },
  clearMarker(id) {
    if (this.marker) {
      this.marker.setMap(null);
      this.marker = null;
    }
    if (this.infoWindow) {
      map.closeInfoWindow(this.infoWindow);
    }
    if (this.markers.length > 0) {
      for (let i = this.markers.length - 1; i >= 0; i--) {
        var marker = this.markers[i];
        if (marker.cid == id) {
          map.removeOverlay(marker);
          marker = null;
          this.markers.splice(i, 1);
          return i;
          // break;
        }
      }
    }
  },
  addLine(options) {
    var points = [];
    options.coords.forEach(data => {
      points.push(new BMap.Point(data[0], data[1]));
    })
    if (this.path) {
      this.clearLine();
    }
    var sy = new BMap.Symbol(BMap_Symbol_SHAPE_BACKWARD_OPEN_ARROW,{
      scale: 0.4, //图标缩放大小
      strokeColor:'#fff', //设置矢量图表的线填充颜色
      strokeWeight: '2', //设置线宽
    });
    var icons = new BMap.IconSequence(sy, '30', '30');
    this.path = new BMap.Polyline(points, {
        strokeColor: "blue",
        strokeWeight: 8, 
        strokeOpacity: 0.8 ,
        icons: [icons]
      });   //创建折线
    map.addOverlay(this.path);

    var view = map.getViewport(eval(points));
    var mapZoom = view.zoom;
    var centerPoint = view.center; map.centerAndZoom(centerPoint, mapZoom);
    map.centerAndZoom(centerPoint, mapZoom)
  
    // var startIcon = new BMap.Icon("https://webapi.amap.com/theme/v1.3/markers/n/start.png", new BMap.Size(19, 31));
    // var endIcon = new BMap.Icon("https://webapi.amap.com/theme/v1.3/markers/n/end.png", new BMap.Size(19, 31));
    // markerStart = new BMap.Marker(points[0], {
    //   icon: startIcon,
    //   offset: new BMap.Size(0, -15)
    // });
    // markerEnd = new BMap.Marker(points[points.length - 1], { 
    //   icon: endIcon, 
    //   offset: new BMap.Size(0, -15) 
    // });
    // console.log(points.length)
    // map.addOverlay(markerStart);
    // map.addOverlay(markerEnd);
  },
  clearLine() {
    if (this.path) {
      map.removeOverlay(this.path);
      this.path = null;
    }

  },
  translate(points) {

  },
  getLocation(options) {
    var geoc = new BMap.Geocoder();
    var pt = new BMap.Point(options.x, options.y);
    geoc.getLocation(pt, function (rs) {
      var address = rs.address;
      options.callback(address);
    });
  }
}