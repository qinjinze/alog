var pi = 3.14159265358979324;
var a = 6378245.0;
var ee = 0.00669342162296594323;

var proj = {
    // 百度坐标转火星坐标
    baiduTomars(bdlonlat) {
        var x_pi = 3.14159265358979324 * 3000.0 / 180.0;
        var mars_point = { lon: 0, lat: 0 };
        var x = bdlonlat[0] - 0.0065;
        var y = bdlonlat[1] - 0.006;
        var z = Math.sqrt(x * x + y * y) - 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) - 0.000003 * Math.cos(x * x_pi);
        var lon = z * Math.cos(theta);
        var lat = z * Math.sin(theta);
        return [lon, lat];
    },
    // 火星坐标转百度坐标
    marsTobaidu(lonlat) {
        var x_pi = 3.14159265358979324 * 3000.0 / 180.0;
        var baidu_point = { lon: 0, lat: 0 };
        var x = lonlat[0];
        var y = lonlat[1];
        var z = Math.sqrt(x * x + y * y) + 0.00002 * Math.sin(y * x_pi);
        var theta = Math.atan2(y, x) + 0.000003 * Math.cos(x * x_pi);
        var lon = z * Math.cos(theta) + 0.0065;
        var lat = z * Math.sin(theta) + 0.006;
        return [lon, lat];
    },
    // 经纬度转墨卡托
    lonlatTomercator(lonlat) {
        var x = lonlat[0] * 20037508.34 / 180;
        var y = Math.log(Math.tan((90 + lonlat[1]) * Math.PI / 360)) / (Math.PI / 180);
        y = y * 20037508.34 / 180;
        return [x, y];
    },
    //墨卡托转经纬度
    mercatorTolonlat(mercator) {
        var x = mercator[0] / 20037508.34 * 180;
        var y = mercator[1] / 20037508.34 * 180;
        y = 180 / Math.PI * (2 * Math.atan(Math.exp(y * Math.PI / 180)) - Math.PI / 2);
        return [x, y];
    },
    // WGS84转火星坐标
    wgsTomars(lonlat) {
        var wgLon = lonlat[0];
        var wgLat = lonlat[1];

        if (this.outOfChina(wgLon, wgLat)) {
            return [wgLon, wgLat];
        }
        var dLat = this.transformLat(wgLon - 105.0, wgLat - 35.0);
        var dLon = this.transformLon(wgLon - 105.0, wgLat - 35.0);
        var radLat = wgLat / 180.0 * pi;
        var magic = Math.sin(radLat);
        magic = 1 - ee * magic * magic;
        var sqrtMagic = Math.sqrt(magic);
        dLat = (dLat * 180.0) / ((a * (1 - ee)) / (magic * sqrtMagic) * pi);
        dLon = (dLon * 180.0) / (a / sqrtMagic * Math.cos(radLat) * pi);
        var lat = wgLat + dLat;
        var lon = wgLon + dLon;
        return [lon, lat]

    },
    // 火星坐标转wgs84
    marsTowgs(lonlat) {
        var lng = lonlat[0];
        var lat = lonlat[1];
        //国外不用转
        if (this.outOfChina(lng, lat)) {
            return [lng, lat]
        }
        var dlat = this.transformLat(lng - 105.0, lat - 35.0)
        var dlng = this.transformLon(lng - 105.0, lat - 35.0)
        var radlat = lat / 180.0 * pi
        var magic = Math.sin(radlat)
        var magic = 1 - ee * magic * magic
        var sqrtmagic = Math.sqrt(magic)
        dlat = (dlat * 180.0) / ((a * (1 - ee)) / (magic * sqrtmagic) * pi)
        dlng = (dlng * 180.0) / (a / sqrtmagic * Math.cos(radlat) * pi)
        var mglat = lat + dlat
        var mglng = lng + dlng
        return [lng * 2 - mglng, lat * 2 - mglat];
    },
    // 判断是否在国外
    outOfChina(lon, lat) {
        if ((lon < 72.004 || lon > 137.8347) && (lat < 0.8293 || lat > 55.8271)) {
            return true;
        } else {
            return false;
        }
    },
    transformLat(x, y) {
        var ret = -100.0 + 2.0 * x + 3.0 * y + 0.2 * y * y + 0.1 * x * y + 0.2 * Math.sqrt(Math.abs(x));
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0;
        ret += (20.0 * Math.sin(y * pi) + 40.0 * Math.sin(y / 3.0 * pi)) * 2.0 / 3.0;
        ret += (160.0 * Math.sin(y / 12.0 * pi) + 320 * Math.sin(y * pi / 30.0)) * 2.0 / 3.0;
        return ret;
    },
    transformLon(x, y) {
        var ret = 300.0 + x + 2.0 * y + 0.1 * x * x + 0.1 * x * y + 0.1 * Math.sqrt(Math.abs(x));
        ret += (20.0 * Math.sin(6.0 * x * pi) + 20.0 * Math.sin(2.0 * x * pi)) * 2.0 / 3.0;
        ret += (20.0 * Math.sin(x * pi) + 40.0 * Math.sin(x / 3.0 * pi)) * 2.0 / 3.0;
        ret += (150.0 * Math.sin(x / 12.0 * pi) + 300.0 * Math.sin(x / 30.0 * pi)) * 2.0 / 3.0;
        return ret;
    }
}
 