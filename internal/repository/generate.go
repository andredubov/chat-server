package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Chats -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Messages -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Participants -o ./mocks/ -s "_minimock.go"
