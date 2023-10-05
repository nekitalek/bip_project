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

public class SecondFactorAuth extends AppCompatActivity {
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

                Intent inten = getIntent();
                String user_id = inten.getStringExtra("user_id");
                String password = inten.getStringExtra("password");

                JSONObject jsonObject = new JSONObject();
                JSONObject eConf = new JSONObject();

                try {
                    eConf.put("user_id", Integer.parseInt(user_id));
                    eConf.put("code", Integer.parseInt(code.getText().toString()));
                    eConf.put("device", "android");
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                try {
                jsonObject.put("e_conf", eConf);
                jsonObject.put("new_password", password);
                } catch (JSONException e) {
                throw new RuntimeException(e);
                }

                String filesDir = inten.getStringExtra("filesDir");

                SaveAndReadFile file = new SaveAndReadFile();
                String token = file.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
                file = null;

                SaveAndReadFile file1 = new SaveAndReadFile();
                String token_cookie = file1.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
                file1 = null;

                Log.d("token_cookie", token_cookie);
                Log.d("token", token);
                Log.d("json", jsonObject.toString());

                Headers headers = new Headers.Builder()
                        .add("Content-Type", "application/json")
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();

                HttpClient.send("https://51.250.24.31:65000/auth/change/password/sec_factor", "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                Toast.makeText(SecondFactorAuth.this, jsonResponse, Toast.LENGTH_SHORT).show();
                            }
                        });
                    }
                });

                Intent intent = new Intent(SecondFactorAuth.this, Profile.class);
                intent.putExtra("filesDir", "/data/user/0/com.example.bip_java/files");
                startActivity(intent);
            }
        });
        Cancel.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(SecondFactorAuth.this, Profile.class);
                startActivity(intent);
            }
        });
    }
}