package com.example.bip_java;

import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.text.TextUtils;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.inputmethod.InputMethodManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.RelativeLayout;
import android.widget.Toast;

import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputEditText;

import org.json.JSONException;
import org.json.JSONObject;

import okhttp3.Headers;

public class SecondAuthWindow extends AppCompatActivity {
    Button Verify, Cancel;
    RelativeLayout root_auth;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_second_auth_window);

        Verify=findViewById(R.id.button);
        Cancel=findViewById(R.id.cancel);
        root_auth=findViewById(R.id.root_auth);

        Verify.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                EditText code = (EditText) findViewById(R.id.verification_code);

                // Скрываем клавиатуру
                InputMethodManager inputMethodManager = (InputMethodManager) getSystemService(Context.INPUT_METHOD_SERVICE);
                View view = getCurrentFocus();
                if (view != null) {
                    inputMethodManager.hideSoftInputFromWindow(view.getWindowToken(), 0);
                }

                if (code.getText().toString().length() != 6) {
                    Snackbar.make(root_auth, "Code must contain 6 digits", Snackbar.LENGTH_SHORT).show();
                    return;
                }

                JSONObject jsonObject = new JSONObject();

                SaveAndReadFile file = new SaveAndReadFile();
                String user_id = file.readStringFromFile("user_id.txt", "/data/user/0/com.example.bip_java/files");
                file = null;

                Log.d("user_id11111111", user_id);
                Log.d("code11111111", code.getText().toString());

                try {
                    jsonObject.put("user_id",  Integer.parseInt(user_id));
                    jsonObject.put("code",  Integer.parseInt(code.getText().toString()));
                    jsonObject.put("device", "android");
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                Log.d("123123123", jsonObject.toString());
                Intent inten = getIntent();
                String filesDir = inten.getStringExtra("filesDir");

                SaveAndReadFile file1 = new SaveAndReadFile();
                String token = file1.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
                file1 = null;

                SaveAndReadFile file2 = new SaveAndReadFile();
                String token_cookie = file2.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
                file2 = null;

                Log.d("token_cookie", token_cookie);
                Log.d("token", token);

                Headers headers = new Headers.Builder()
                        .add("Content-Type", "application/json")
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();
                String url = "https://51.250.24.31:65000/auth/sign-up/sec_factor";

                if (inten.getStringExtra("transition").equals("1")) {
                    url = "https://51.250.24.31:65000/auth/sign-in/sec_factor";
                }

                HttpClient.send(url, "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        Log.d("22222", jsonResponse);

                        if (inten.getStringExtra("transition").equals("1")) {
                            int startIndex = jsonResponse.indexOf("\"auth_token\":\"") + "\"auth_token\":\"".length();
                            if (startIndex >= 0) {
                                int endIndex = jsonResponse.indexOf("\"", startIndex);
                                String authToken = jsonResponse.substring(startIndex, endIndex);

                                SaveAndReadFile file = new SaveAndReadFile();
                                file.saveStringToFile("authToken.txt", authToken, getFilesDir().toString());
                                file = null;
                                Log.d("authToken", authToken);
                                Intent intent = new Intent(SecondAuthWindow.this, MainWindowApp.class);
                                startActivity(intent);
                            }
                        }
                        if (inten.getStringExtra("transition").equals("0")) {
                            Intent intent = new Intent(SecondAuthWindow.this, MainActivity.class);
                            startActivity(intent);
                        }
                    }
                });
            }
        });
        Cancel.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(SecondAuthWindow.this, MainActivity.class);
                startActivity(intent);
            }
        });
    }
}