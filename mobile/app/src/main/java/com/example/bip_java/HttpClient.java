package com.example.bip_java;

import android.content.Context;
import android.util.Log;
import android.widget.Toast;

import org.json.JSONObject;
import org.json.JSONException;

import java.io.File;
import java.io.FileWriter;
import java.io.IOException;
import java.net.CookieManager;
import java.net.CookiePolicy;
import java.security.KeyManagementException;
import java.security.NoSuchAlgorithmException;
import java.security.cert.X509Certificate;
import java.util.List;

import okhttp3.CookieJar;
import java.net.CookieManager;
import java.net.CookiePolicy;

import okhttp3.*;

import com.example.bip_java.JsonResponseCallback;
import com.example.bip_java.MainActivity;

import javax.net.ssl.HostnameVerifier;
import javax.net.ssl.HttpsURLConnection;
import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSession;
import javax.net.ssl.TrustManager;
import javax.net.ssl.X509TrustManager;

public class HttpClient {
    public static void send(String url, String typeRequest, JSONObject jsonObject, Headers headers, boolean flag, final JsonResponseCallback callback) {

        TrustManager[] trustAllCertificates = new TrustManager[]{
                new X509TrustManager() {
                    public X509Certificate[] getAcceptedIssuers() {
                        return new X509Certificate[0];
                    }

                    public void checkClientTrusted(X509Certificate[] certs, String authType) {
                    }

                    public void checkServerTrusted(X509Certificate[] certs, String authType) {
                    }
                }
        };

        SSLContext sslContext = null;
        try {
            sslContext = SSLContext.getInstance("TLS");
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException(e);
        }
        try {
            sslContext.init(null, trustAllCertificates, new java.security.SecureRandom());
        } catch (KeyManagementException e) {
            throw new RuntimeException(e);
        }


        OkHttpClient client = new OkHttpClient.Builder()
                .sslSocketFactory(sslContext.getSocketFactory(), (X509TrustManager) trustAllCertificates[0])
                .hostnameVerifier((hostname, session) -> true)
                .build();

        Request request = null;

        if (typeRequest.equals("post")) {

            MediaType JSON = MediaType.parse("application/json; charset=utf-8");
            RequestBody requestBody = RequestBody.create(JSON, String.valueOf(jsonObject));

            request = new Request.Builder()
                .url(url)
                .post(requestBody)
                .headers(headers)
                .build();
        }
        else if (typeRequest.equals("get")){
            if (headers == null){
                request = new Request.Builder()
                    .url(url)
                    .get()
                    .build();
            }
            else {
                request = new Request.Builder()
                    .url(url)
                    .get()
                    .headers(headers)
                    .build();
            }
        }
        else if (typeRequest.equals("del")){

            MediaType JSON = MediaType.parse("application/json; charset=utf-8");
            RequestBody requestBody = RequestBody.create(JSON, String.valueOf(jsonObject));

            request = new Request.Builder()
                .url(url)
                .delete(requestBody)
                .headers(headers)
                .build();
        }


        client.newCall(request).enqueue(new Callback() {
            @Override
            public void onFailure(Call call, IOException e) {
                // Обработка ошибки
                callback.onJsonResponse(e.toString());
            }

            @Override
            public void onResponse(Call call, Response response) throws IOException {
                if (response.isSuccessful()) {
                    // Получение ответа в виде JSON файла
                    ResponseBody responseBody = response.body();
                    if (responseBody != null) {
                        String jsonResponse = responseBody.string();

                        if ("GET".equalsIgnoreCase(response.request().method())) {
                            if (flag == true) {
                                String cookie = response.headers("Set-Cookie").toString();

                                if (cookie != null && !cookie.isEmpty()) {
                                    String result = cookie.substring(cookie.indexOf("[") + 1, cookie.indexOf("]"));
                                    Log.d("asdasdasdasdasdasd", result);

                                    SaveAndReadFile file = new SaveAndReadFile();
                                    file.saveStringToFile("token_cookie.txt", result, "/data/user/0/com.example.bip_java/files");
                                    file = null;

                                }
                            }
                        }
                        callback.onJsonResponse(jsonResponse);

                        response.close();
                    }
                } else {
                    // Обработка ошибки
                    int httpStatusCode = response.code(); // Код состояния HTTP
                    String errorMessage = response.message(); // Сообщение об ошибке
                    callback.onJsonResponse("HTTP Error: " + httpStatusCode + " - " + errorMessage);
                }
                client.dispatcher().executorService().shutdown();
                client.connectionPool().evictAll();
            }
        });
    }
}
