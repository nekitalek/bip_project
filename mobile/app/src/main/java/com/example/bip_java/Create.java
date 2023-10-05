package com.example.bip_java;

import static android.widget.Toast.LENGTH_SHORT;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.webkit.CookieManager;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.core.content.ContextCompat;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import java.util.ArrayList;
import java.util.List;

import com.example.bip_java.MyAdapter;
import com.example.bip_java.JsonResponseCallback;
import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputEditText;

import org.json.JSONException;
import org.json.JSONObject;

import okhttp3.Headers;


public class Create extends AppCompatActivity {
    Button Create, Back;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.create);

        Create = findViewById(R.id.create);
        Back = findViewById(R.id.back_create);

        Create.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                TextInputEditText Sports = findViewById(R.id.nameSports);
                TextInputEditText Date = findViewById(R.id.nameDate);
                TextInputEditText StartTime = findViewById(R.id.nameStartTime);
                TextInputEditText EndTime = findViewById(R.id.nameEndTime);
                TextInputEditText Address = findViewById(R.id.nameAddress);
                TextInputEditText Description = findViewById(R.id.nameDescription);

                if (Sports.getText().toString().length() == 0 || Date.getText().toString().length() == 0 ||
                    StartTime.getText().toString().length() == 0 || EndTime.getText().toString().length() == 0 ||
                    Address.getText().toString().length() == 0) {

                    Toast.makeText(Create.this, "Fill in all the fields", Toast.LENGTH_SHORT).show();
                    return;
                }

                JSONObject jsonObject = new JSONObject();

                try {
                    jsonObject.put("time_start", Date.getText().toString() + "T" + StartTime.getText().toString() + "Z");
                    jsonObject.put("time_end", Date.getText().toString() + "T" + EndTime.getText().toString() + "Z");
                    jsonObject.put("place", Address.getText().toString());
                    jsonObject.put("game", Sports.getText().toString());
                    jsonObject.put("description", Description.getText().toString());
                    jsonObject.put("public", false);
                } catch (JSONException e) {
                    throw new RuntimeException(e);
                }
                Log.d("111111111", jsonObject.toString());
                Intent inten = getIntent();
                String filesDir = inten.getStringExtra("filesDir");

                SaveAndReadFile file = new SaveAndReadFile();
                String authToken = file.readStringFromFile("authToken.txt", "/data/user/0/com.example.bip_java/files");
                file = null;

                SaveAndReadFile file1 = new SaveAndReadFile();
                String token = file1.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
                file1 = null;

                SaveAndReadFile file2 = new SaveAndReadFile();
                String token_cookie = file2.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
                file2 = null;

                Log.d("authToken", authToken);
                Log.d("token_cookie", token_cookie);
                Log.d("token", token);
                Headers headers = new Headers.Builder()
                        .add("Authorization", "Bearer" + " " + authToken)
                        .add("Content-Type", "application/json")
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();

                HttpClient.send("https://51.250.24.31:65000/api/event", "post", jsonObject, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        Log.d("event", jsonResponse);
                        Intent intent = new Intent(Create.this, MainWindowApp.class);
                        startActivity(intent);
                    }
                });


            }
        });

        Back.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(Create.this, MainWindowApp.class);
                startActivity(intent);
            }
        });
    }
}
