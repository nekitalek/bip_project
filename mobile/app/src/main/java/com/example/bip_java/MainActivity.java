package com.example.bip_java;

import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;

import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.text.TextUtils;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.widget.Button;
import android.widget.RelativeLayout;
import android.widget.Toast;

import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputEditText;

import java.io.File;
import java.io.InputStream;
import java.security.KeyStore;
import java.security.SecureRandom;
import java.security.cert.Certificate;
import java.security.cert.CertificateException;
import java.security.cert.CertificateFactory;
import java.security.cert.X509Certificate;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.http.conn.ssl.SSLSocketFactory;
import org.json.JSONException;
import org.json.JSONObject;

import com.example.bip_java.SaveAndReadFile;
import com.google.gson.Gson;

import javax.net.ssl.HostnameVerifier;
import javax.net.ssl.HttpsURLConnection;
import javax.net.ssl.SSLContext;
import javax.net.ssl.SSLSession;
import javax.net.ssl.X509TrustManager;

import okhttp3.Headers;
import okhttp3.OkHttpClient;

public class MainActivity extends AppCompatActivity {
    RelativeLayout root;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button btnSignIn = findViewById(R.id.btnSignIn);
        Button btnRegister = findViewById(R.id.btnRegister);
        root = findViewById(R.id.root_element);

        btnRegister.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) { showRegisterWindow(); }
        });
        btnSignIn.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                showSignInWindow();
            }
        });
    }
    private void showSignInWindow() {

        AlertDialog.Builder dialog = new AlertDialog.Builder(this);
        dialog.setTitle("Sign In");
        dialog.setMessage("Enter all data for login");

        LayoutInflater inflater = LayoutInflater.from(this);
        View sign_in_window = inflater.inflate(R.layout.sign_in_window, null);
        dialog.setView(sign_in_window);

        TextInputEditText email = sign_in_window.findViewById(R.id.emailField);
        TextInputEditText pass = sign_in_window.findViewById(R.id.passField);

        dialog.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialogInterface, int which) {
                dialogInterface.dismiss();
            }
        });

        dialog.setPositiveButton("Login", new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialogInterface, int which) {
                if (TextUtils.isEmpty(email.getText().toString())) {
                    Snackbar.make(root, "Empty email!", Snackbar.LENGTH_SHORT).show();
                    return;
                }
                if (pass.getText().toString().length() < 5) {
                    Snackbar.make(root, "Short password!(min 5 symbols)", Snackbar.LENGTH_SHORT).show();
                    return;
                }

                HttpClient.send("https://51.250.24.31:65000/CSRF", "get", null, null, true, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        int startIndex = jsonResponse.indexOf("\"token_CSRF\":\"") + "\"token_CSRF\":\"".length();
                        int endIndex = jsonResponse.indexOf("\"}", startIndex);
                        String tokenValue = jsonResponse.substring(startIndex, endIndex);
                        Log.d("tokenValue", tokenValue);
                        SaveAndReadFile file = new SaveAndReadFile();
                        file.saveStringToFile("token_CSRF.txt", tokenValue, getFilesDir().toString());
                        file = null;

                    }
                });

                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }

                JSONObject jsonObject = new JSONObject();

                try {
                    jsonObject.put("Login", email.getText().toString());
                    jsonObject.put("Password", pass.getText().toString());
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                SaveAndReadFile file = new SaveAndReadFile();
                String token = file.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
                file = null;

                SaveAndReadFile file2 = new SaveAndReadFile();
                String token_cookie = file2.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
                file2 = null;

                Log.d("token_cookie", token_cookie);
                Log.d("token", token);
                Log.d("json", jsonObject.toString());

                Headers headers = new Headers.Builder()
                        .add("Content-Type", "application/json")
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();
                HttpClient.send("https://51.250.24.31:65000/auth/sign-in/password", "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        Log.d("post_auth", jsonResponse);
                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                Toast.makeText(MainActivity.this, jsonResponse, Toast.LENGTH_SHORT).show();
                            }
                        });
                    }
                });

                Intent intent = new Intent(MainActivity.this, SecondAuthWindow.class);
                intent.putExtra("transition", "1");
                intent.putExtra("filesDir", getFilesDir().toString());
                startActivity(intent);
            }
        });
        dialog.show();

    }
    private void showRegisterWindow() {
        AlertDialog.Builder dialog = new AlertDialog.Builder(this);
        dialog.setTitle("Registration");
        dialog.setMessage("Enter all data for registration");

        LayoutInflater inflater = LayoutInflater.from(this);
        View register_window = inflater.inflate(R.layout.register_windows, null);
        dialog.setView(register_window);

        TextInputEditText email = register_window.findViewById(R.id.emailField);
        TextInputEditText name = register_window.findViewById(R.id.nameField);
        TextInputEditText pass = register_window.findViewById(R.id.passField);

        dialog.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialogInterface, int which) {
                dialogInterface.dismiss();
            }
        });

        dialog.setPositiveButton("Add", new DialogInterface.OnClickListener() {
            @Override
            public void onClick(DialogInterface dialogInterface, int which) {

                if(TextUtils.isEmpty(email.getText().toString())) {
                    Snackbar.make(root, "Empty email!", Snackbar.LENGTH_SHORT).show();
                    return;
                }
                if(TextUtils.isEmpty(name.getText().toString())) {
                    Snackbar.make(root, "Empty name!", Snackbar.LENGTH_SHORT).show();
                    return;
                }
                if(pass.getText().toString().length() < 5) {
                    Snackbar.make(root, "Short password!(min 5 symbols)", Snackbar.LENGTH_SHORT).show();
                    return;
                }
                HttpClient.send("https://51.250.24.31:65000/CSRF", "get", null, null, true, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        int startIndex = jsonResponse.indexOf("\"token_CSRF\":\"") + "\"token_CSRF\":\"".length();
                        int endIndex = jsonResponse.indexOf("\"}", startIndex);
                        String tokenValue = jsonResponse.substring(startIndex, endIndex);
                        Log.d("tokenValue", tokenValue);
                        SaveAndReadFile file = new SaveAndReadFile();
                        file.saveStringToFile("token_CSRF.txt", tokenValue, getFilesDir().toString());
                        file = null;

                    }
                });

                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }

                JSONObject jsonObject = new JSONObject();

                try {
                    jsonObject.put("Login", email.getText().toString());
                    jsonObject.put("Username", name.getText().toString());
                    jsonObject.put("Password", pass.getText().toString());
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                SaveAndReadFile file = new SaveAndReadFile();
                String token = file.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
                file = null;
                Log.d("token", token);

                SaveAndReadFile file1 = new SaveAndReadFile();
                String token_cookie = file1.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
                file1 = null;

                Log.d("token_cookie", token_cookie);
                Headers headers = new Headers.Builder()
                        .add("Content-Type", "application/json")
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();

                HttpClient.send("https://51.250.24.31:65000/auth/sign-up/password", "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        Log.d("11111", jsonResponse);

                        Pattern pattern = Pattern.compile("\"user_id\":(\\d+)");
                        Matcher matcher = pattern.matcher(jsonResponse);
                        String user_id = "0";
                        if (matcher.find()) {
                            user_id = matcher.group(1);
                        }

                        SaveAndReadFile file = new SaveAndReadFile();
                        file.saveStringToFile("user_id.txt", user_id, getFilesDir().toString());
                        file = null;
                    }
                });

                Intent intent = new Intent(MainActivity.this, SecondAuthWindow.class);
                intent.putExtra("transition", "0");
                intent.putExtra("filesDir", getFilesDir().toString());
                startActivity(intent);

            }
        });
        dialog.show();
    }
}