package com.example.bip_java;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;
import androidx.core.content.ContextCompat;
import androidx.recyclerview.widget.LinearLayoutManager;
import androidx.recyclerview.widget.RecyclerView;

import java.lang.reflect.Type;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import com.example.bip_java.MyAdapter;
import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import org.json.JSONException;
import org.json.JSONObject;

import okhttp3.Headers;

public class MainWindowApp extends AppCompatActivity implements MyAdapter.OnItemClickListener {
    Button Profile, Logout, Create, Update;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.main_window_app);

        Profile = findViewById(R.id.profile);
        Logout = findViewById(R.id.logout);
        Create = findViewById(R.id.buttonCreate);
        Update = findViewById(R.id.update);
        RecyclerView recyclerView = findViewById(R.id.recyclerView);


        Profile.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent inten = getIntent();
                Intent intent = new Intent(MainWindowApp.this, Profile.class);
                intent.putExtra("filesDir", inten.getStringExtra("filesDir"));
                startActivity(intent);
            }
        });

        Create.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent inten = getIntent();
                Intent intent = new Intent(MainWindowApp.this, Create.class);
                intent.putExtra("filesDir", inten.getStringExtra("filesDir"));
                startActivity(intent);
            }
        });

        Logout.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(MainWindowApp.this, MainActivity.class);
                startActivity(intent);
            }
        });

        Update.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
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
                        .add("Cookie", token_cookie + ';')
                        .add("X-CSRF-TOKEN", token)
                        .build();

                try {
                    Thread.sleep(1000);
                } catch (InterruptedException e) {
                    throw new RuntimeException(e);
                }

                HttpClient.send("https://51.250.24.31:65000/api/event", "get", null, headers, false, new JsonResponseCallback() {
                    @Override
                    public void onJsonResponse(String jsonResponse) {
                        Log.d("event", jsonResponse);
                        SaveAndReadFile file = new SaveAndReadFile();
                        file.saveStringToFile("view.txt", jsonResponse, "/data/user/0/com.example.bip_java/files");
                        file = null;
                    }
                });
                SaveAndReadFile file3 = new SaveAndReadFile();
                String view = file3.readStringFromFile("view.txt", "/data/user/0/com.example.bip_java/files");
                file3 = null;

                Log.d("zzzzzz", view);

                if (view != null) {
                    Gson gson = new Gson();
                    Type eventItemType = new TypeToken<List<EventItem>>() {
                    }.getType();

                    List<EventItem> eventItems = gson.fromJson(view, eventItemType);

                    if (eventItems != null) {
                        List<String> data1 = new ArrayList<>();
                        List<String> data2 = new ArrayList<>();
                        MyAdapter adapter = new MyAdapter(data1, data2, MainWindowApp.this);
                        recyclerView.setAdapter(adapter);
                        recyclerView.setLayoutManager(new LinearLayoutManager(MainWindowApp.this));


                        for (EventItem eventItem : eventItems) {
                            String str1 = eventItem.getGame() + " " + eventItem.getTimeStart().substring(0, 10) + " " +
                                    eventItem.getTimeStart().substring(11, 16) + "-" + eventItem.getTimeEnd().substring(11, 16);
                            String str2 = eventItem.getPlace() + " " + eventItem.getDescription() + " " + "id:" + eventItem.getEventId();
                            Log.d("str1", str1);
                            Log.d("str2", str2);
                            data1.add(str1);
                            data2.add(str2);
                        }
                        adapter.notifyDataSetChanged();
                    }
                }
            }
        });
    }

    @Override
    public void onItemClick(int position, TextView sport, TextView address, Button buttonRecord) {
        Toast.makeText(this, sport.getText().toString() + address.getText().toString(), Toast.LENGTH_SHORT).show();
        if (buttonRecord.getText().toString().equals("join")) {
            SaveAndReadFile file = new SaveAndReadFile();
            String user_id = file.readStringFromFile("user_id.txt", "/data/user/0/com.example.bip_java/files");
            file = null;

            SaveAndReadFile file1 = new SaveAndReadFile();
            String view = file1.readStringFromFile("view.txt", "/data/user/0/com.example.bip_java/files");
            file1 = null;

            SaveAndReadFile file2 = new SaveAndReadFile();
            String authToken = file2.readStringFromFile("authToken.txt", "/data/user/0/com.example.bip_java/files");
            file2 = null;

            SaveAndReadFile file3 = new SaveAndReadFile();
            String token = file3.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
            file3 = null;

            SaveAndReadFile file4 = new SaveAndReadFile();
            String token_cookie = file4.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
            file4 = null;

            JSONObject jsonObject = new JSONObject();

            Pattern pattern = Pattern.compile("id:(\\d+)");
            Matcher matcher = pattern.matcher(address.getText().toString());

            int event_id = 0;

            if (matcher.find()) {
                String idValue = matcher.group(1);
                event_id = Integer.parseInt(idValue);
            }

            try {
                jsonObject.put("event_id",  event_id);
                jsonObject.put("user_id",  Integer.parseInt(user_id));
                jsonObject.put("status", "Confirmed");
            } catch (JSONException e) {
                throw new RuntimeException(e);
            }

            Log.d("authToken", authToken);
            Log.d("token_cookie", token_cookie);
            Log.d("token", token);
            Headers headers = new Headers.Builder()
                    .add("Authorization", "Bearer" + " " + authToken)
                    .add("Cookie", token_cookie + ';')
                    .add("Content-Type", "application/json")
                    .add("X-CSRF-TOKEN", token)
                    .build();

            HttpClient.send("https://51.250.24.31:65000/api/invitation", "post", jsonObject, headers, false, new JsonResponseCallback() {
                @Override
                public void onJsonResponse(String jsonResponse) {
                    Log.d("post_invitation", jsonResponse);
                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            Toast.makeText(MainWindowApp.this, jsonResponse, Toast.LENGTH_SHORT).show();
                        }
                    });
                }
            });

            buttonRecord.setText("disconnect");
            buttonRecord.setBackgroundColor(ContextCompat.getColor(this, R.color.red));
        }
        else if (buttonRecord.getText().toString().equals("disconnect")){
            SaveAndReadFile file = new SaveAndReadFile();
            String user_id = file.readStringFromFile("user_id.txt", "/data/user/0/com.example.bip_java/files");
            file = null;

            SaveAndReadFile file1 = new SaveAndReadFile();
            String view = file1.readStringFromFile("view.txt", "/data/user/0/com.example.bip_java/files");
            file1 = null;

            SaveAndReadFile file2 = new SaveAndReadFile();
            String authToken = file2.readStringFromFile("authToken.txt", "/data/user/0/com.example.bip_java/files");
            file2 = null;

            SaveAndReadFile file3 = new SaveAndReadFile();
            String token = file3.readStringFromFile("token_CSRF.txt", "/data/user/0/com.example.bip_java/files");
            file3 = null;

            SaveAndReadFile file4 = new SaveAndReadFile();
            String token_cookie = file4.readStringFromFile("token_cookie.txt", "/data/user/0/com.example.bip_java/files");
            file4 = null;

            JSONObject jsonObject = new JSONObject();

            Pattern pattern = Pattern.compile("id:(\\d+)");
            Matcher matcher = pattern.matcher(address.getText().toString());

            int event_id = 0;

            if (matcher.find()) {
                String idValue = matcher.group(1);
                event_id = Integer.parseInt(idValue);
            }

            try {
                jsonObject.put("user_id", Integer.parseInt(user_id));
                jsonObject.put("event_id", event_id);
            } catch (JSONException e) {
                throw new RuntimeException(e);
            }

            Log.d("authToken", authToken);
            Log.d("token_cookie", token_cookie);
            Log.d("token", token);

            Headers headers = new Headers.Builder()
                    .add("Authorization", "Bearer" + " " + authToken)
                    .add("Cookie", token_cookie + ';')
                    .add("Content-Type", "application/json")
                    .add("X-CSRF-TOKEN", token)
                    .build();


            HttpClient.send("https://51.250.24.31:65000/api/invitation", "del", jsonObject, headers, false, new JsonResponseCallback() {
                @Override
                public void onJsonResponse(String jsonResponse) {
                    Log.d("del_invitation", jsonResponse);
                    runOnUiThread(new Runnable() {
                        @Override
                        public void run() {
                            Toast.makeText(MainWindowApp.this, jsonResponse, Toast.LENGTH_SHORT).show();
                        }
                    });
                }
            });
            buttonRecord.setText("join");
            buttonRecord.setBackgroundColor(ContextCompat.getColor(this, R.color.green));
        }
    }
}
