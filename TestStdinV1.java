package com.container.s;

import com.github.dockerjava.api.DockerClient;
import com.github.dockerjava.api.command.ExecCreateCmdResponse;
import com.github.dockerjava.core.DefaultDockerClientConfig;
import com.github.dockerjava.core.DockerClientBuilder;
import com.github.dockerjava.core.DockerClientConfig;
import com.github.dockerjava.core.command.ExecStartResultCallback;
import com.github.dockerjava.netty.NettyDockerCmdExecFactory;
import java.io.ByteArrayInputStream;

public class TestStdinV1 {
    public static void main(String[] args) throws InterruptedException {
        // 创建 Docker 客户端
        DockerClientConfig config = DefaultDockerClientConfig.createDefaultConfigBuilder()
                .withDockerHost("tcp://10.24.2.232:2375").build();
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
        ExecStartResultCallback callback = new ExecStartResultCallback(System.out, System.err);

        //String input = "ls / \n";
        //ByteArrayInputStream inputStream = new ByteArrayInputStream(input.getBytes());

        System.out.println("输入命令:");

       // System.setIn(inputStream);
        dockerClient.execStartCmd(execCreateCmdResponse.getId()).withStdIn(System.in)
                .exec(callback).awaitCompletion();
    }
}
