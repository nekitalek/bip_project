package com.example.bip_java;

import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.text.TextUtils;
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

public class ChangeEmail extends AppCompatActivity {
    Button Verify, Cancel;
    RelativeLayout root_auth;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.second_auth_change_email);

        Verify=findViewById(R.id.button_change_email);
        Cancel=findViewById(R.id.cancel_change_email);
        root_auth=findViewById(R.id.root_change_email);

        Verify.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                EditText CodeOldEmail = (EditText) findViewById(R.id.old_email);
                EditText CodeNewEmail = (EditText) findViewById(R.id.new_email);

                // Скрываем клавиатуру
                InputMethodManager inputMethodManager = (InputMethodManager) getSystemService(Context.INPUT_METHOD_SERVICE);
                View view = getCurrentFocus();
                if (view != null) {
                    inputMethodManager.hideSoftInputFromWindow(view.getWindowToken(), 0);
                }

                if (CodeOldEmail.getText().toString().length() != 6) {
                    Snackbar.make(root_auth, "Code must contain 6 digits", Snackbar.LENGTH_SHORT).show();
                    return;
                }
                if (CodeNewEmail.getText().toString().length() != 6) {
                    Snackbar.make(root_auth, "Code must contain 6 digits", Snackbar.LENGTH_SHORT).show();
                    return;
                }

                Intent inten = getIntent();
                String user_id = inten.getStringExtra("user_id");
                String new_login = inten.getStringExtra("new_login");

                JSONObject jsonObject = new JSONObject();
                JSONObject eConf = new JSONObject();

                try {
                    eConf.put("user_id", user_id);
                    eConf.put("code", CodeOldEmail.getText().toString());
                    eConf.put("device", "Android");
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                try {
                    jsonObject.put("e_conf", eConf);
                    jsonObject.put("new_login", new_login);
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }

                String filesDir = inten.getStringExtra("filesDir");

                SaveAndReadFile file = new SaveAndReadFile();
                String token = file.readStringFromFile("token.txt", filesDir);
                file = null;

                Headers headers = new Headers.Builder()
                        .add("Content-Type", "application/json")
                        .add("X-CSRF-TOKEN", token)
                        .build();

                HttpClient.send("https://reqres.in/api/register", "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                Toast.makeText(ChangeEmail.this, jsonResponse, Toast.LENGTH_SHORT).show();
                            }
                        });
                    }
                });

                JSONObject jsonObject2 = new JSONObject();

                try {
                    jsonObject2.put("user_id", user_id);
                    jsonObject2.put("code", CodeNewEmail.getText().toString());
                    jsonObject2.put("device", "Android");
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }

                HttpClient.send("https://reqres.in/api/register", "post", jsonObject2, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        runOnUiThread(new Runnable() {
                            @Override
                            public void run() {
                                Toast.makeText(ChangeEmail.this, jsonResponse, Toast.LENGTH_SHORT).show();
                            }
                        });
                    }
                });

                Intent intent = new Intent(ChangeEmail.this, Profile.class);
                startActivity(intent);
            }
        });
        Cancel.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(ChangeEmail.this, Profile.class);
                startActivity(intent);
            }
        });
    }
}