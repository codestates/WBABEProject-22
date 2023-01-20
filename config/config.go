package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server struct {
		Mode string
		Port string
	}

	DB map[string]string

	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}
}

func GetConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
		}
	}

	return c
}
/* [코드리뷰]
	 * 해당 코드에는 하나의 function에서 간결한 이중 조건문이 발생하고 있습니다.
	 * 15 line으로 구성된 한눈에 들어오는 함수에서는 사용하기 적합합니다.
	 * 그러나 코드 라인수가 많아지고, 비즈니스 로직이 풍부해 지는 것을 고려한다면
	 * return으로 나갈 수 빠지는 case를 최소화하고, if문을 줄이는 방향으로 개발이 진행되어야 합니다. 
	 * 점진적으로 코드의 가독성이 조금 더 향상하게 됩니다.
	 *
	 * 또한 하나의 function에서 예상되는 기능이 여러가지로 분류되는 것을 볼 수 있습니다.
	 * return config 뿐만 아니라, panic이라는 내장함수를 사용하여 프로그램을 종료시키는 코드도 담고 있습니다.
	 * 해당 기능을 config만을 return 시키되, nil 값을 return 시키는 코드로 통일성을 주고.
	 * GetConfig 함수를 호출하는 코드에서 if c == nil 인 case를 예외처리 해주는 방식으로 코딩을 해보시는 것을 추천드립니다.
	 * 이를 통해 만약 오류가 발생될 때, 어디에서 오류가 발생했는지를 일관된 흐름 안에서 파악이 가능해질 것으로 예상됩니다.
	 * as-is: 
	 if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
		}
	}
	 * to-be:
	 if file, err := os.Open(fpath); err == nil {
			defer file.Close()
			//toml 파일 디코딩
			if err := toml.NewDecoder(file).Decode(c); err == nil {
				fmt.Println(c)
				return c, nil
			} 
		}
		return nil, err
	 */


// References
// Class material: lecture 12