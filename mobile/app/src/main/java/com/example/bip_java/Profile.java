package com.example.bip_java;

import android.content.DialogInterface;
import android.content.Intent;
import android.os.Bundle;
import android.text.InputType;
import android.text.TextUtils;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.webkit.CookieManager;
import android.widget.Button;
import android.widget.EditText;
import android.widget.ImageButton;
import android.widget.RelativeLayout;
import android.widget.Toast;

import androidx.appcompat.app.AlertDialog;
import androidx.appcompat.app.AppCompatActivity;

import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputEditText;
import com.google.gson.Gson;

import org.json.JSONException;
import org.json.JSONObject;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

import okhttp3.Headers;

public class Profile extends AppCompatActivity {

    private EditText passwordEditText;
    private ImageButton showPasswordButton;

    private boolean isPasswordVisible = false;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.profile_settings);

        RelativeLayout root_set = findViewById(R.id.root_settings);
        TextInputEditText name = findViewById(R.id.nameProfile);
        TextInputEditText email = findViewById(R.id.emailProfile);
        Button LogoutAllDevices = findViewById(R.id.logoutAllDevices);

        SaveAndReadFile file = new SaveAndReadFile();
        String authToken = file.readStringFromFile("authToken.txt", "/data/user/0/com.example.bip_java/files");
        file = null;

        SaveAndReadFile file1 = new SaveAndReadFile();
        String token = file1.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
        file1 = null;

        SaveAndReadFile file2 = new SaveAndReadFile();
        String token_cookie = file2.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
        file2 = null;

        SaveAndReadFile file3 = new SaveAndReadFile();
        String user_id = file3.readStringFromFile("user_id.txt", "/data/user/0/com.example.bip_java/files");
        file3 = null;

        Log.d("authToken", authToken);
        Log.d("token_cookie", token_cookie);
        Log.d("token", token);
        Headers headers = new Headers.Builder()
                .add("Authorization", "Bearer" + " " + authToken)
                .add("Cookie", token_cookie + ';')
                .add("X-CSRF-TOKEN", token)
                .build();

        String url = "https://51.250.24.31:65000/api/user/" + user_id;
        Log.d("user_id", user_id);

        HttpClient.send(url, "get", null, headers, false, new JsonResponseCallback() {
            @Override
            public void onJsonResponse(String jsonResponse) {
                Log.d("api/user/1", jsonResponse);
                SaveAndReadFile file = new SaveAndReadFile();
                file.saveStringToFile("info.txt", jsonResponse, "/data/user/0/com.example.bip_java/files");
                file = null;
            }
        });

        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            throw new RuntimeException(e);
        }

        SaveAndReadFile file4 = new SaveAndReadFile();
        String info = file4.readStringFromFile("info.txt", "/data/user/0/com.example.bip_java/files");
        file4 = null;

        Gson gson = new Gson();
        User user = gson.fromJson(info, User.class);

        name.setText(user.getUsername());
        email.setText(user.getLogin());

        name.setFocusable(false);
        name.setEnabled(false);
        email.setFocusable(false);
        email.setEnabled(false);

        /*Button chUsr = findViewById(R.id.change_username);
        Button chPsw = findViewById(R.id.change_password);
        Button chMail = findViewById(R.id.change_email);

        chUsr.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                AlertDialog.Builder dialog = new AlertDialog.Builder(Profile.this);
                dialog.setTitle("Change username");

                LayoutInflater inflater = LayoutInflater.from(Profile.this);
                View change_usr = inflater.inflate(R.layout.change_username, null);
                dialog.setView(change_usr);

                TextInputEditText newUsr = change_usr.findViewById(R.id.newUsr);

                dialog.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        dialogInterface.dismiss();
                    }
                });

                dialog.setPositiveButton("Change", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        if(newUsr.getText().toString().equals(name.getText().toString())) {
                            Snackbar.make(root_set, "New username matches with old username", Snackbar.LENGTH_SHORT).show();
                            return;
                        }

                        // Отправить на сервер новое имя

                        Intent intent = new Intent(Profile.this, SecondFactorAuth.class);
                        startActivity(intent);
                    }
                });
                dialog.show();
            }
        });

        chPsw.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                AlertDialog.Builder dialog = new AlertDialog.Builder(Profile.this);
                dialog.setTitle("Change password");

                LayoutInflater inflater = LayoutInflater.from(Profile.this);
                View change_pass = inflater.inflate(R.layout.change_password, null);
                dialog.setView(change_pass);

                TextInputEditText oldPsw = change_pass.findViewById(R.id.oldPsw);
                TextInputEditText newPsw = change_pass.findViewById(R.id.newPsw);
                TextInputEditText newPswRep = change_pass.findViewById(R.id.newPswRep);

                dialog.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        dialogInterface.dismiss();
                    }
                });

                dialog.setPositiveButton("Change", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        if(newPsw.getText().toString().equals(newPswRep.getText().toString())) {
                            if(newPsw.getText().toString().length() < 5) {
                                Snackbar.make(root_set, "Short password!(min 5 symbols)", Snackbar.LENGTH_SHORT).show();
                                return;
                            }
                        }
                        else{
                            Snackbar.make(root_set, "Passwords mismatch", Snackbar.LENGTH_SHORT).show();
                            return;
                        }

                        HttpClient.send("https://158.160.27.251:443/CSRF", "get", null, null, true, new JsonResponseCallback() {
                            @Override
                            public void onJsonResponse(String jsonResponse) {
                                int startIndex = jsonResponse.indexOf("\"token_CSRF\":\"") + "\"token_CSRF\":\"".length();
                                int endIndex = jsonResponse.indexOf("\"}", startIndex);
                                String tokenValue = jsonResponse.substring(startIndex, endIndex);
                                Log.d("tokenValue", tokenValue);
                                SaveAndReadFile file = new SaveAndReadFile();
                                file.saveStringToFile("token_CSRF.txt", tokenValue, "/data/user/0/com.example.bip_java/files");
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
                            jsonObject.put("Login", "inanmasov2017@mail.ru");
                            jsonObject.put("Password", oldPsw.getText().toString());
                        } catch (JSONException e) {
                            throw new RuntimeException(e);
                        }

                        Intent inten = getIntent();
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

                        HttpClient.send("https://158.160.27.251:443/auth/change/password/first_factor", "post", jsonObject, headers, false, new JsonResponseCallback() {
                            @Override
                            public void onJsonResponse(String jsonResponse) {
                                Log.d("chpass", jsonResponse);
                                Pattern pattern = Pattern.compile("\"user_id\":(\\d+)");
                                Matcher matcher = pattern.matcher(jsonResponse);
                                String user_id = "0";
                                if (matcher.find()) {
                                    user_id = matcher.group(1);
                                }

                                SaveAndReadFile file = new SaveAndReadFile();
                                file.saveStringToFile("user_id_chpsw.txt", user_id, "/data/user/0/com.example.bip_java/files");
                                file = null;
                                runOnUiThread(new Runnable() {
                                    @Override
                                    public void run() {
                                        Toast.makeText(Profile.this, jsonResponse, Toast.LENGTH_SHORT).show();
                                    }
                                });
                            }
                        });

                        SaveAndReadFile file2 = new SaveAndReadFile();
                        String user_id = file2.readStringFromFile("user_id_chpsw.txt", getFilesDir().toString());
                        file2 = null;

                        Intent intent = new Intent(Profile.this, SecondFactorAuth.class);
                        intent.putExtra("user_id", user_id);
                        intent.putExtra("password", newPsw.getText().toString());
                        startActivity(intent);
                    }
                });
                dialog.show();
            }
        });

        chMail.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                AlertDialog.Builder dialog = new AlertDialog.Builder(Profile.this);
                dialog.setTitle("Change email");

                LayoutInflater inflater = LayoutInflater.from(Profile.this);
                View change_email = inflater.inflate(R.layout.change_email, null);
                dialog.setView(change_email);

                TextInputEditText newEmail = change_email.findViewById(R.id.newEmail);

                dialog.setNegativeButton("Cancel", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        dialogInterface.dismiss();
                    }
                });

                dialog.setPositiveButton("Change", new DialogInterface.OnClickListener() {
                    @Override
                    public void onClick(DialogInterface dialogInterface, int which) {
                        if(newEmail.getText().toString().equals(email.getText().toString())) {
                            Snackbar.make(root_set, "New email matches with old email", Snackbar.LENGTH_SHORT).show();
                            return;
                        }

                        JSONObject jsonObject = new JSONObject();

                        try {
                            //jsonObject.put("Login", email.getText().toString());
                            //jsonObject.put("Password", password.getText().toString());
                            jsonObject.put("email", "eve.holt@reqres.in");
                            jsonObject.put("password", "pistol");
                        } catch (JSONException e) {
                            throw new RuntimeException(e);
                        }

                        Intent inten = getIntent();
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
                                        Toast.makeText(Profile.this, jsonResponse, Toast.LENGTH_SHORT).show();
                                    }
                                });
                            }
                        });

                        Intent intent = new Intent(Profile.this, ChangeEmail.class);
                        intent.putExtra("user_id", "1");
                        intent.putExtra("new_login", newEmail.getText().toString());
                        startActivity(intent);
                    }
                });
                dialog.show();
            }
        });*/

        Button back = findViewById(R.id.back);
        back.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(Profile.this, MainWindowApp.class);
                startActivity(intent);
            }
        });

        LogoutAllDevices.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(Profile.this, MainActivity.class);
                startActivity(intent);
            }
        });
    }
}
