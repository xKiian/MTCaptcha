function hashStr(_0x333e63) {
    var _0x4b487d = 0x0;
    var _0x1f5117;
    if (_0x333e63.length == 0x0) {
        return _0x4b487d;
    }
    for (_0x1f5117 = 0x0; _0x1f5117 < _0x333e63.length; _0x1f5117++) {
        var _0x4f6494 = _0x333e63.charCodeAt(_0x1f5117);
        _0x4b487d = (_0x4b487d << 0x5) - _0x4b487d + _0x4f6494;
        _0x4b487d = _0x4b487d & _0x4b487d;
    }
    return _0x4b487d;
}

const URLSafeBase64Int2CharMap = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', '-', '_']

function URLSafeBase64IntToChar(_0xc8a0bf) {
    if (0x0 > _0xc8a0bf || _0xc8a0bf > 0x3f) {
        throw "arg i must be between 0 .. 63 inclusive";
    }
    return URLSafeBase64Int2CharMap[_0xc8a0bf % 0x40];
}


function URLSafeBase4096IntToChar(_0x227456) {
    if (_0x227456 > 0xfff || _0x227456 < 0x0) {
        throw "arg i must be between 0 .. 4095 inclusive";
    }
    return '' + URLSafeBase64IntToChar(_0x227456 >> 0x6) + URLSafeBase64IntToChar(_0x227456 & 0x3f);
}


_0x5eb8fb = {
    "v": [0, 1],
    "r": [3675, 1423],
    "n": 1750027007,
    "z": -12,
    "a": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
    "d": "https://service.mtcaptcha.com/mtcv1/client/iframe.html?v=2024-11-14.21.25.03&sitekey=MTPublic-KzqLY1cKH&iframeId=mtcaptcha-iframe-1&widgetSize=standard&custom=false&widgetInstance=mtcaptcha&challengeType=standard&theme=basic&lang=en&action=&autoFadeOuterText=false&host=https%3A%2F%2F2captcha.com&hostname=2captcha.com&serviceDomain=service.mtcaptcha.com&textLength=0&lowFrictionInvisible=&enableMouseFlow=false&resetTS=1750025115193",
    "l": "en-US",
    "h": 10
}
var _0x1fc26c = [0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0];
_0x5eb8fb.res = _0x1fc26c;
_0x1fc26c[0x0] = _0x5eb8fb.v[0x0];
_0x1fc26c[0x1] = _0x5eb8fb.v[0x1];
_0x1fc26c[0x2] = _0x5eb8fb.r[0x0];
_0x1fc26c[0x3] = _0x5eb8fb.r[0x1];
_0x379605: {
    var _0xfa34f3 = Math.floor(_0x5eb8fb.n / 0x400000) % 0x1000;
    var _0x172215 = (_0x5eb8fb.n % 0x400000 >> 0xb) % 0x1000;
    var _0x8201d4 = _0x5eb8fb.n % 0x400000 % 0x800;
    _0x1fc26c[0x4] = _0xfa34f3;
    _0x1fc26c[0x5] = _0x172215;
    _0x1fc26c[0x6] = _0x8201d4;
}
_0x1fc26c[0x7] = Math.abs(Math.floor(_0x5eb8fb.z / 0xa) % 0x1000);
_0x1fc26c[0x8] = Math.abs(hashStr((_0x5eb8fb.a + '').toLowerCase())) % 0x1000;
var _0x2d86bb = (_0x5eb8fb.d + '').toLowerCase().match("^(?:https?://)?(?:[^@/\n]+@)?(?:www.)?([^:/\n]+)");
_0x2d86bb = _0x2d86bb == null ? '' : _0x2d86bb[0x1];
_0x1fc26c[0x9] = Math.abs(hashStr(_0x2d86bb.toLowerCase())) % 0x1000;
_0x1fc26c[0xa] = Math.abs(hashStr((_0x5eb8fb.l + '').toLowerCase())) % 0x1000;
_0x1fc26c[0xb] = _0x5eb8fb.h % 0x1000;
var _0xaddae8 = 0x0;
for (var _0x1ad7be = 0x0; _0x1ad7be < 0xc; _0x1ad7be++) {
    _0xaddae8 = _0xaddae8 * 0x1f + _0x1fc26c[_0x1ad7be];
    _0xaddae8 = _0xaddae8 & _0xaddae8;
}
_0xaddae8 = Math.abs(_0xaddae8);
_0x1fc26c[0xc] = _0xaddae8 % 0x1000;
for (var _0x1ad7be = 0x4; _0x1ad7be < _0x1fc26c.length; _0x1ad7be++) {
    _0x1fc26c[_0x1ad7be] = _0x1fc26c[_0x1ad7be] ^ _0x5eb8fb.r[_0x1ad7be % 0x2];
}
var _0x11126a = [];
for (var _0x1ad7be = 0x0; _0x1ad7be < _0x1fc26c.length; _0x1ad7be++) {
    _0x11126a.push(URLSafeBase4096IntToChar(_0x1fc26c[_0x1ad7be]));
}

console.log(_0x11126a.join(""))