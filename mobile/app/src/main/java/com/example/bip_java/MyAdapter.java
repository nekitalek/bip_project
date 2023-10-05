package com.example.bip_java;

import android.annotation.SuppressLint;
import android.content.Context;
import android.util.Log;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.Button;
import android.widget.TextView;

import androidx.core.content.ContextCompat;
import androidx.recyclerview.widget.RecyclerView;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import okhttp3.Headers;

public class MyAdapter extends RecyclerView.Adapter<MyAdapter.ViewHolder> {
    private List<String> data1, data2;

    public interface OnItemClickListener {
        void onItemClick(int position, TextView textView1, TextView textView2, Button button_record);
    }
    private OnItemClickListener listener;

    public MyAdapter(List<String> data1, List<String> data2, OnItemClickListener listener) {
        this.data1 = data1;
        this.data2 = data2;
        this.listener = listener;
    }

    @Override
    public ViewHolder onCreateViewHolder(ViewGroup parent, int viewType) {
        Context context = parent.getContext();
        LayoutInflater inflater = LayoutInflater.from(context);
        View view = inflater.inflate(R.layout.list_item, parent, false);
        return new ViewHolder(view);
    }

    @Override
    public void onBindViewHolder(ViewHolder holder, @SuppressLint("RecyclerView") int position) {
        String item1 = data1.get(position);
        String item2 = data2.get(position);
        holder.textView1.setText(item1);
        holder.textView2.setText(item2);

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

        Pattern pattern = Pattern.compile("id:(\\d+)");
        Matcher matcher = pattern.matcher(item2);

        int event_id = 0;

        if (matcher.find()) {
            String idValue = matcher.group(1);
            event_id = Integer.parseInt(idValue);
        }

        String url = "https://51.250.24.31:65000/api/event?event_items_id=" + event_id;

        HttpClient.send(url, "get", null, headers, false, new JsonResponseCallback() {
            @Override
            public void onJsonResponse(String jsonResponse) {
                Log.d("event?event_items_id", jsonResponse);

                SaveAndReadFile file = new SaveAndReadFile();
                String user_id = file.readStringFromFile("user_id.txt", "/data/user/0/com.example.bip_java/files");
                file = null;

                Gson gson = new Gson();
                Type listType = new TypeToken<List<EventItem>>() {}.getType();
                List<EventItem> eventItems = gson.fromJson(jsonResponse, listType);

                for (EventItem eventItem : eventItems) {
                    if (eventItem.getParticipantsWithUserId(Integer.parseInt(user_id)) == true) {
                        holder.button_record.setText("disconnect");
                        holder.button_record.setBackgroundColor(ContextCompat.getColor(holder.itemView.getContext(), R.color.red));
                    } else {
                        holder.button_record.setText("join");
                        holder.button_record.setBackgroundColor(ContextCompat.getColor(holder.itemView.getContext(), R.color.green));
                    }
                }
            }
        });


        holder.button_record.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                if (listener != null) {
                    listener.onItemClick(position, holder.textView1, holder.textView2, holder.button_record);
                }
            }
        });
    }

    @Override
    public int getItemCount() {
        return data1.size();
    }

    public class ViewHolder extends RecyclerView.ViewHolder {
        public TextView textView1, textView2;
        public Button button_record;

        public ViewHolder(View itemView) {
            super(itemView);
            textView1 = itemView.findViewById(R.id.sport);
            textView2 = itemView.findViewById(R.id.address);
            button_record = itemView.findViewById(R.id.button_record);
        }
    }
}