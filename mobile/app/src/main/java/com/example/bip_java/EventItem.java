package com.example.bip_java;

import com.google.gson.annotations.SerializedName;

import java.util.ArrayList;
import java.util.List;

public class EventItem {
    @SerializedName("event_items_id")
    private int eventItemsId;
    private int admin;
    private List<Participant> participant;
    @SerializedName("time_start")
    private String timeStart;
    @SerializedName("time_end")
    private String timeEnd;
    private String place;
    private String game;
    private String description;
    @SerializedName("public")
    private boolean publ;

    public List<Participant> getParticipant() { return participant; }

    public boolean getParticipantsWithUserId(int userId) {
        List<Participant> participantsWithUserId = new ArrayList<>();
        for (Participant participant : participant) {
            if (participant.getUserId() == userId) {
                return true;
            }
            else{
                return false;
            }
        }
        return false;
    }
    public String getTimeStart() {
        return timeStart;
    }

    public String getTimeEnd() {
        return timeEnd;
    }

    public String getPlace() {
        return place;
    }

    public String getGame() {
        return game;
    }

    public String getDescription() {
        return description;
    }

    public int getEventId() {
        return eventItemsId;
    }
}
