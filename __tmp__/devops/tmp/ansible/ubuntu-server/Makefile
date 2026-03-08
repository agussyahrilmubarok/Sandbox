ansible.playbook.ssh.key.used:
	ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/$(file) --key-file./id_rsa

ansible.playbook.ssh.key:
	SSHPASSWORD=$(passwd) USER_SSH_KEY="$(cat ./id_rsa.pub)" ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/$(file)

ansible.playbook.ssh:
	SSHPASSWORD=$(passwd) ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/$(file)

ansible.playbook:
	ansible-playbook -i inventory.yaml --vault-password-file .vault-password playbooks/$(file)

ansible.vault:
	ansible-vault encrypt_string --vault-password-file .vault-password "$(pwgen 16 1)" --name "password"

vault.passwd:
	pwgen 16 1 > .vault-password

test.server:
	ping $(SERVER)

passwd.random:
	openssl rand --hex 10

ssh.key:
	ssh-keygen -t rsa -b 4096 -f ./id_rsa