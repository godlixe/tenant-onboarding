package services

// type AuthService struct {
// 	userRepository repository.UserRepository
// }

// func NewAuthService(
// 	userRepository repository.UserRepository,
// ) *AuthService {
// 	return &AuthService{
// 		userRepository: userRepository,
// 	}
// }

// // Login returns an access token if successful
// // and returns an error otherwise.
// func (s *AuthService) Login(
// 	ctx context.Context,
// 	userRequest *entity.User,
// ) (string, error) {
// 	user, err := s.userRepository.FindByUsername(ctx, userRequest.Username)
// 	if err != nil {
// 		return "", err
// 	}

// 	isPasswordMatch, err := auth.ComparePassword(
// 		user.Password,
// 		[]byte(userRequest.Password),
// 	)
// 	if err != nil {
// 		return "", err
// 	}

// 	if !isPasswordMatch {
// 		return "", errors.New("password does not match")
// 	}

// 	token, err := auth.GenerateJWTToken(
// 		user.ID.String(),
// 	)
// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil
// }

// func (s *AuthService) Register(
// 	ctx context.Context,
// 	userRequest *entity.User,
// ) (*entity.User, error) {
// 	_, err := s.userRepository.FindByUsername(ctx, userRequest.Username)
// 	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound)) {
// 		return nil, err
// 	}

// 	if err == nil {
// 		return nil, errors.New("user with same username exists")
// 	}

// 	err = s.userRepository.CreateUser(ctx, userRequest)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return userRequest, nil
// }

// func (s *AuthService) Me(
// 	ctx context.Context,
// ) (*entity.User, error) {
// 	var userId uuid.UUID

// 	id := ctx.Value("user_id").(string)
// 	userId = uuid.MustParse(id)
// 	res, err := s.userRepository.FindById(ctx, userId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return res, nil
// }
