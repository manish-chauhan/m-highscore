import (
	pbhighscore "github.com/manish-chauhan/m-apis.git"
	"google.golang.org/grpc"
	"github.com/rs/zerolog/log"
)

type Grpc struct {
	address string
	srv *grpc.Server
}

var HighScore  = 999999999.0

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse , error){
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
	Set: true,
	}, nil

}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
	HighScore: HighScore,
	}, nil
}

