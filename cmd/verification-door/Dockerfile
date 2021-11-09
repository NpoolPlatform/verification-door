FROM jrei/systemd-centos:7

COPY VerificationDoor.viper.yaml /etc/VerificationDoor/VerificationDoor.viper.yaml
COPY verification-door /usr/local/bin/verification-door

CMD ["verification-door", "run"]
