package com.example.bip_java;

import android.content.Context;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.OutputStreamWriter;
import com.example.bip_java.MainActivity;
public class SaveAndReadFile {
    public void saveStringToFile(String fileName, String data, String filesDir) {
        try {
            File file = new File(filesDir, fileName);
            FileOutputStream fos = new FileOutputStream(file, false);
            OutputStreamWriter outputStreamWriter = new OutputStreamWriter(fos);

            outputStreamWriter.write(data);
            outputStreamWriter.close();

            fos.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
    public String readStringFromFile(String fileName, String filesDir) {
        StringBuilder stringBuilder = new StringBuilder();
        try {
            File file = new File(filesDir, fileName);
            FileInputStream fis = new FileInputStream(file);
            InputStreamReader inputStreamReader = new InputStreamReader(fis);
            BufferedReader bufferedReader = new BufferedReader(inputStreamReader);

            String line;
            while ((line = bufferedReader.readLine()) != null) {
                stringBuilder.append(line);
            }

            bufferedReader.close();
            inputStreamReader.close();
            fis.close();
        } catch (IOException e) {
            e.printStackTrace();
        }

        return stringBuilder.toString();
    }
}
