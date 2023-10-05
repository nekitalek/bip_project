package com.example.bip_java;

import com.google.gson.annotations.SerializedName;

public class User {
    @SerializedName("user_id")
    private int userId;

    private String login;
    private String username;

    public int getUserId() {
        return userId;
    }
    public String getLogin() {
        return login;
    }
    public String getUsername() {
        return username;
    }
}