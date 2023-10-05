package com.example.bip_java;

import com.google.gson.annotations.SerializedName;

public class Participant {
    @SerializedName("user_id")
    private int userId;
    private String username;

    public int getUserId() {
        return userId;
    }
}