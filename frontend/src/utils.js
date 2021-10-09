import router from "@/router";

export const calculatingPage = (currPage, totalPage) => {
    var reserveLength = 5;
    var marginRight = 2;
    var marginLeft = 2;
    var pages = [];
    if (currPage < reserveLength) {
        for(var i=1; i<=reserveLength;i++) {
            if (i > totalPage) {
                break;
            }
            pages.push(i);
        }
    } else {
        if (currPage+marginRight > totalPage) {
            marginLeft += ((currPage+marginRight) - totalPage);
        }
        for(var i=currPage-marginLeft; i<=currPage+marginLeft; i++) {
            if (i > totalPage) {
                break;
            }
            pages.push(i);
        }
    }
    return pages;
}

export const universalTimeString = (n) => {
    var splitted = n.split(" ");
    let dateNewFormat = splitted[0].replace(/-/g, "/")+" "+splitted[1]+" +0700";
    return dateNewFormat
}

export const isFloat = (n) => {
    var regexp = /^\d+(\.\d{1,3})?$/;
    if (regexp.test(n) === true) {
        return true;
    }
    return false;
}

export const isNotFloat = (n) => {
    var regexp = /^\d+$/;
    if (regexp.test(n) === true) {
        return true;
    }
    return false;
}

export const calculatePriceNumber = (qty) => {
    if (Number.isNaN(parseInt(qty))) {
        return "-";
    }
    return qty*1;
};

export const numberWithCommas = (x) => {
    return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

export const sleep = (ms) => {
    return new Promise(resolve => setTimeout(resolve, ms));
}

export const countryPrice = (c, p) => {
  if (c === "IDN") {
      return "Rp. "+p;
  } else {
      return "Eur. "+p;
  }
};

export const logout = () => {
    deleteCookie("islogged");
    router.push({name: 'login', query: {logout: true}});
};

export const setCookie = (cname, cvalue, exdays) => {
    var d = new Date();
    d.setTime(d.getTime() + (exdays * 24 * 60 * 60 * 1000));
    var expires = "expires="+d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

export const deleteCookie = (cname, exdays) => {
    setCookie(cname, "", -1000);
}

export const getCookie = (cname) => {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for(var i = 0; i <ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
};

export const isLoggedIn = () => {
    var islogged = getCookie("islogged");
    if (islogged === "1") {
        return true;
    }
    return false;
};

export const getBaseUrl = () => {
    let hostVar = window.location.hostname;
    let protocol = window.location.protocol;
    if (window.location.port) {
        if (hostVar === 'localhost' || hostVar === '127.0.0.1') {
            hostVar = protocol+`//${hostVar}:8080/`;
        } else {
            // hostVar = `${hostVar}:${window.location.port}`;
            hostVar = protocol+`//${hostVar}:8080/`;
        }
    } else {
        hostVar = protocol+`//${hostVar}/api/`;
    }
    return hostVar;
};


export const getRegisterUrl = () => {
    var hostVar = window.location.hostname;
    let protocol = window.location.protocol;
    if (process.env.NODE_ENV === 'production') {
        return protocol+"//"+hostVar+"/user/api-register";
    }
    return protocol+"//"+hostVar+":8001"+"/user/api-register";
};

export const getForgotPasswordUrl = () => {
    var hostVar = window.location.hostname;
    let protocol = window.location.protocol;
    if (process.env.NODE_ENV === 'production') {
        return protocol+"//"+hostVar+"/user/request-password-reset";
    }
    return protocol+"//"+hostVar+":8001"+"/user/request-password-reset";
};

export const encryptText = (p) => {
    var pidCrypt = require("pidcrypt");
    var pidCryptUtil = require("pidcrypt/pidcrypt_util");
    require("pidcrypt/rsa");

    var rsa = new pidCrypt.RSA();
    var modulus = "wZyLn4YzpJHd0w8jdfGM1IkqCA4y6W0sDe2lOzkiEXgLK/Pu8DLxFPJZI35hnoCRSzd5a6OdzcTx/ZGn1Xc4Qamc7nv+6ehhmXWOdQzAIiO1t8jkyOr6sSrHtlviHfd8CZlshz6lfwWcJF9tIBJI4ew0ZiMikXdnPEOOrwDZsQyGrAObDpDtERedHs2TddEfzC3hc49iVt15XRoAwmAg6cfFGklPXMoJN7bYI8m2/TJsXo7jknQLYzqxOzRLh9SpoHqpPdAoE9YCiqybA8js9SlxTnhUlhcXz5wjZxdx17ydoOApcGAzbs0dgQKFbXQUKJ9KpL6ibduECF8/dX3TCx5Q//bF13tlbBlr3tzOqt21WCflpKO3R5+p8RQylEkNMTR8IJ9rpcF3f3f8aAWMg48gBH3QyvC1QNhWGZVcf2E8PvBFpoyZRDe+6r8sSGDrq+o1oOI9OHrCaHMKNwXxkQMJc1LmBg/YhkWBXUktCQ0Z5T5oyIKpjCBxDSWkMB2V1s3K6XHQWI2c41jeWet12o4s+RufT69tpl9zncSykfIA5YnMETSPlFgBqsT+DKrQ/yXTEPgDhy0VqBu3GzI7MGHn4uPEHerJP5b+pSsB/9UM5JNdtd5u2pHKmD3kaFyQbkf261xtH0u65ogLhxB9ui4zBHPPBFiu0QJml/VGaNc=";
    var e = 'AQAB';

    var n = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(modulus));
    e = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(e));

    rsa.setPublic (n, e, 16);
    var encryptedPassword = rsa.encryptRaw(p);
    var data = pidCryptUtil.encodeBase64(pidCryptUtil.convertFromHex(encryptedPassword));
    return data;
};


// export const decryptPassword = (p) => {
//     var pidCrypt = require("pidcrypt");
//     var pidCryptUtil = require("pidcrypt/pidcrypt_util");
//     require("pidcrypt/rsa");
//
//     var rsa = new pidCrypt.RSA();
//
//     var n = 'wZyLn4YzpJHd0w8jdfGM1IkqCA4y6W0sDe2lOzkiEXgLK/Pu8DLxFPJZI35hnoCRSzd5a6OdzcTx/ZGn1Xc4Qamc7nv+6ehhmXWOdQzAIiO1t8jkyOr6sSrHtlviHfd8CZlshz6lfwWcJF9tIBJI4ew0ZiMikXdnPEOOrwDZsQyGrAObDpDtERedHs2TddEfzC3hc49iVt15XRoAwmAg6cfFGklPXMoJN7bYI8m2/TJsXo7jknQLYzqxOzRLh9SpoHqpPdAoE9YCiqybA8js9SlxTnhUlhcXz5wjZxdx17ydoOApcGAzbs0dgQKFbXQUKJ9KpL6ibduECF8/dX3TCx5Q//bF13tlbBlr3tzOqt21WCflpKO3R5+p8RQylEkNMTR8IJ9rpcF3f3f8aAWMg48gBH3QyvC1QNhWGZVcf2E8PvBFpoyZRDe+6r8sSGDrq+o1oOI9OHrCaHMKNwXxkQMJc1LmBg/YhkWBXUktCQ0Z5T5oyIKpjCBxDSWkMB2V1s3K6XHQWI2c41jeWet12o4s+RufT69tpl9zncSykfIA5YnMETSPlFgBqsT+DKrQ/yXTEPgDhy0VqBu3GzI7MGHn4uPEHerJP5b+pSsB/9UM5JNdtd5u2pHKmD3kaFyQbkf261xtH0u65ogLhxB9ui4zBHPPBFiu0QJml/VGaNc=';
//     var e = 'AQABAAAAAAA=';
//     var d = 'RPXsaWqUlGZ8O0PmHoll3wHoHaxpB+mhymg99SU5dD1sxUACeD28zvgvNcw7GwiyN7dPoT5K49LlWqr5u7fWKtUuF2fE+S2TLeTIU/qxqLdNb2O30bSYjGgvcE2z8XaIEKc8F8QBkIfAANiYdtbsbMQEOekD1ApSlA2AEedcoopSoYw7O2n1dqQRi7ovbsSI7eHdrpqArW3/X2D2qkUxE/Paya0Y/nTb6aNPwxeC/RZrsRZ2CmEY+UQbgnny+TOPI6DtDR5KWB5fE2O9ZxC3De+v8wI7JqadDY6QsZSRMQV2o6mUOPma+DlI6/2o5r50u69ZF+rq0fRhR9YdNU0W0Pw6FTpDe4w40fGGSDEjyfyKdC4xQO/duPzykASpb+/BhWpoe3A7AP9ntaycbe88jJ5is4O/7EAb47YOtGtjOPGWxnfuHfk6Jfec/EP+x/sxUlB2G/wQqXZuLnTi/S3gAiiiyu6E5DvtvaRcbQw/d5qAHhnItFBgorztqtnxcGW4HPpfmR0drvQeR664m3QVD4E1RMMzHVKoDN4d6KKn7t6kO+Z7osCo0AXBWNXdxmmFgEZV2uDjgWJqExz/wf3bYsuCj2q3dcqvzc7zjbzBi6GJNFTBqvTwjJFLblYIcdzAH5bpLX+jBDqxa38yzL4hSG6e9Neo9v6mjPkzFTu1oik=';
//
//     n = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(n));
//     d = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(d));
//     e = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(e));
//
//     rsa.setPrivate(n,e,d,16);
//     console.log(p);
//     p = pidCryptUtil.convertToHex(pidCryptUtil.decodeBase64(p));
//
//     var decryptedPassword = rsa.decryptRaw(p);
//     return decryptedPassword;
// };
