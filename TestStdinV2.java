package com.container.s;

import com.github.dockerjava.api.DockerClient;
import com.github.dockerjava.api.async.ResultCallback;
import com.github.dockerjava.api.command.ExecCreateCmdResponse;
import com.github.dockerjava.api.command.ExecStartCmd;
import com.github.dockerjava.api.model.Frame;
import com.github.dockerjava.core.DefaultDockerClientConfig;
import com.github.dockerjava.core.DockerClientBuilder;
import com.github.dockerjava.core.DockerClientConfig;
import com.github.dockerjava.core.command.ExecStartResultCallback;
import com.github.dockerjava.netty.NettyDockerCmdExecFactory;

import java.io.*;
import java.util.Scanner;

public class TestStdinV2 {
    public static void main(String[] args) throws InterruptedException, IOException {
        // 创建 Docker 客户端
        DockerClientConfig config = DefaultDockerClientConfig.createDefaultConfigBuilder()
                .withDockerHost("tcp://10.24.2.232:2375").withApiVersion("1.43").build();
        DockerClient dockerClient = DockerClientBuilder.getInstance(config)
                .withDockerCmdExecFactory(new NettyDockerCmdExecFactory()).build();

        String containerId = "b6f1524cd741";

        // 执行 Bash 命令以启动交互会话
        ExecCreateCmdResponse execCreateCmdResponse = dockerClient.execCreateCmd(containerId)
                .withAttachStdout(true)
                .withAttachStderr(true)
                .withAttachStdin(true)
                .withTty(true)
                .withCmd("bash")
                .exec();

        // 开始执行并保持交互会话
        String input = "cat  /tmp/abc.txt \n";

        PipedOutputStream outputStream = new PipedOutputStream();
        PipedInputStream inputStream = new PipedInputStream(outputStream);
        PipedOutputStream result = new PipedOutputStream();

        //outputStream.connect(inputStream);

        ExecStartResultCallback callback = new ExecStartResultCallback(){
            @Override
            public void onNext(Frame frame) {
                if (frame != null) {
                    String res=new String(frame.getPayload());
                    System.out.println(res);

                }
            }
        };

        Thread t1=new Thread(new Runnable() {
            public void run() {

                Scanner scanner = new Scanner(System.in);
                while (true){
                    try {
                        System.out.println("输入命令:");
                        outputStream.write((scanner.nextLine()+" \n").getBytes());
                        outputStream.flush();
                    } catch (IOException e) {
                        throw new RuntimeException(e);
                    }
                }
            }
        });
        t1.start();
/*
        Thread t2=new Thread(new Runnable() {
            public void run() {
                Scanner scanner = new Scanner(System.in);
                byte[] cache=new byte[1024];
                while (true){
                    try {
                        inputStream.read(cache);
                        System.out.println("读取内容:");
                        System.out.println(new String(cache).trim());
                    } catch (IOException e) {
                        throw new RuntimeException(e);
                    }
                }
            }
        });
        t2.start(); */

        ExecStartCmd execStartCmd =dockerClient.execStartCmd(execCreateCmdResponse.getId());
           execStartCmd.withStdIn(inputStream);
            ResultCallback resultCallback= execStartCmd.exec(callback).awaitCompletion();
            System.out.println("##################");

    }
}
