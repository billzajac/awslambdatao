package main

// FIXME: Finish adding passages 28-81

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var isbnRegexp = regexp.MustCompile(`[0-9]{3}\-[0-9]{10}`)
var errorLogger = log.New(os.Stderr, "ERROR ", log.Llongfile)
var debugLogger = log.New(os.Stderr, "DEBUG ", log.Llongfile)

func router(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return show(req)
	default:
		return clientError(http.StatusMethodNotAllowed)
	}
}

func show(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	index, passage := GetPassage()
	//fmt.Fprintf(w, "%d. %s", index, passage)

	//js, err := json.Marshal(bk)
	//if err != nil {
	//	return serverError(err)
	//}
	//	Body:       string(js),

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("%d. %s", index, passage),
	}, nil
}

func serverError(err error) (events.APIGatewayProxyResponse, error) {
	errorLogger.Println(err.Error())

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       http.StatusText(http.StatusInternalServerError),
	}, nil
}

func clientError(status int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       http.StatusText(status),
	}, nil
}

func main() {
	lambda.Start(router)
}

func GetPassage() (int, string) {
	var passages [81]string
	passages[0] = `THE BEGINNING OF POWER
The Tao that can be expressed
is not the Tao of the Absolute. The name that can be named
Is not the name of the Absolute.
The nameless originated Heaven and Earth. The named is the Mother of All Things.
Thus, without expectation,
One will always perceive the subtlety; And, with expectation,
One will always perceive the boundary.
The source of these two is identical, Yet their names are different. Together they are called profound, Profound and mysterious, The gateway to the Collective Subtlety.
`
	passages[1] = `USING POLARITY
When all the world knows beauty as beauty,
There is ugliness. When they know good as good,
There there is evil.
In this way
Existence and nonexistence produce each other. Difficult and easy complete each other. Long and short contrast each other. High and low attract each other. Pitch and tone harmonize each other. Future and past follow each other.
Therefore, Evolved Individuals
Hold their position without effort, Practice their philosophy without words, Are a part of All Things and overlook nothing. They produce but do not possess, Act without expectation, Succeed without taking credit.
Since, indeed, they take no credit, it remains with them.
`
	passages[2] = `KEEPING PEACE
Do not exalt the very gifted,
And people will not contend. Do not treasure goods that are hard to get,
And people will not become thieves. Do not focus on desires,
And people's minds will not be confused.
Therefore, Evolved Individuals lead oti is by
Opening their minds, Reinforcing their centers, Relaxing their desires, Strengthening their characters.
Let the people always act without strategy or desire;
Let the clever not venture to act. Act without action,
And nothing is without order.
`
	passages[3] = `THE NATURE OF THE TAO
The Tao is empty and yet useful; Somehow it never fills up. So profound! It resembles the source of All Things.
It blunts the sharpness, Unties the tangles, And harmonizes the brightness. It identifies with the ways of the world.
So deep! It resembles a certain existence. I do not know whose offspring it is, This Image in front of the source.
`
	passages[4] = `HOLDING TO THE CENTER
Heaven and Earth are impartial;
They regard All Things as straw dogs. Evolved Individuals are impartial;
They regard all people as straw dogs.
Between Heaven and Earth,
The space is like a bellows. The shape changes,
But not the form. The more it moves,
The more it produces.
Too much talk will exhaust itself. It is better to remain centered.
`
	passages[5] = `PERCEIVING THE SUBTLE
The mystery of the valley is immortal;
It is known as the Subtle Female. The gateway of the Subtle Female
Is the source of Heaven and Earth.
Everlasting, endless, it appears to exist. Its usefulness comes with no effort.
`
	passages[6] = `THE POWER OF SELFLESSNESS
Heaven is eternal, the Earth everlasting. They can be eternal and everlasting Because they do not exist for themselves. For that reason they can exist eternally.
Therefore, Evolved Individuals
Put themselves last,
And yet they are first. Put themselves outside,
And yet they remain.
Is it not because they are without self-interest That their interests succeed?
`
	passages[7] = `NONCOMPETITIVE VALUES
The highest value is like water.
The value in water benefits All Things,
And yet it does not contend. It stays in places that others disdain,
And therefore is close to the Tao.
The value in a dwelling is location. The value in a mind is depth. The value in relations is benevolence. The value in words is sincerity. The value in leadership is order. The value in work is competence. The value in effort is timeliness.
Since, indeed, they do not contend, There is no resentment.
`
	passages[8] = `TRANSCENDING DECLINE
Holding to fullness Is not as good as stopping in time.
Sharpness that probes Cannot protect for long.
A house filled with riches Cannot be defended.
Pride in wealth and position Is overlooking one's collapse.
Withdrawing when success is achieved
Is the Tao in Nature.
`
	passages[9] = `INNER HARMONY
In managing your instincts and embracing Oneness,
Can you be undivided? In focusing your Influence,
Can you yield as a newborn child? In clearing your insight,
Can you become free of error? In loving people and leading the organization,
Can you take no action? In opening and closing the gateway to nature,
Can you not weaken? In seeing clearly in all directions,
Can you be without knowledge?

Produce things, cultivate things;
Produce but do not possess.
Act without expectation.
Advance without dominating.
These are called the Subtle Powers.
`
	passages[10] = `USING WHAT IS NOT
Thirty spokes converge at one hub;
What is not there makes the wheel useful. Clay is shaped to form a vessel;
What is not there makes the vessel useful. Doors and windows are cut to form a room;
What is not there makes the room useful.
Therefore, take advantage of what is there, By making use of what is not.
`
	passages[11] = `CONTROLLING THE SENSES
The five colors will blind one's eye. The five tones will deafen one's ear. The five flavors will jade one's taste.
Racing and hunting will derange one's mind. Goods that are hard to get will obstruct one's way.
Therefore, Evolved Individuals Regard the center and not the eye. Hence they discard one and receive the other.
`
	passages[12] = `EXPANDING IDENTIFICATION
There is alarm in both favor and disgrace. Esteem and fear are identified with the self.
What is the meaning of "alarm in both favor and disgrace?" Favor ascends; disgrace descends. To attain them brings alarm. To lose them brings alarm. That is the meaning of "alarm in both favor and disgrace."
What is the meaning of "esteem and fear are identified with the self?" The reason for our fear Is the presence of our self. When we are selfless, What is there to fear?
Therefore those who esteem the world as self
Will be committed to the world. Those who love the world as self
Will be entrusted with the world.
`
	passages[13] = `THE ESSENCE OF TAO
Looked at but not seen:
Its name is formless. Listened to but not heard:
Its name is soundless. Reached for but not obtained:
Its name is intangible.
These three cannot be analyzed, So they mingle and act as one.
Its rising is not bright;
Its setting is not dark. Endlessly, the nameless goes on,
Merging and returning to nothingness.
That is why it is called
The form of the formless,
The image of nothingness. That is why it is called elusive.
Confronted, its beginning is not seen. Followed, its end is not seen.
Hold on to the ancient Tao;
Control the current reality. Be aware of the ancient origins;
This is called the Essence of Tao.
`
	passages[14] = `THE POWER IN SUBTLE FORCE
Those skillful in the ancient Tao Are subtly ingenious and profoundly intuitive. They are so deep they cannot be recognized. Since, indeed, they cannot be recognized, Their force can be contained.

So careful!
As if wading a stream in winter. So hesitant!
As if respecting all sides in the community. So reserved!
As if acting as a guest. So yielding!
As if ice about to melt. So candid!
As if acting with simplicity. So open!
As if acting as a valley. So integrated!
As if acting as muddy water.
Who can harmonize with muddy water,
And gradually arrive at clarity? Who can move with stability,
And gradually bring endurance to life?
Those who maintain the Tao
Do not desire to become full. Indeed, since they are not full,
They can be used up and also renewed.
`
	passages[15] = `KNOWING THE ABSOLUTE
Attain the highest openness;
Maintain the deepest harmony. Become a part of All Things;
In this way, I perceive the cycles.
Indeed, things are numerous;
But each cycle merges with the source. Merging with the source is called harmonizing;
This is known as the cycle of destiny.
The cycle of destiny is called the Absolute;
Knowing the Absolute is called insight. To not know the Absolute
Is to recklessly become a part of misfortune.
To know the Absolute is to be tolerant.
What is tolerant becomes impartial; What is impartial becomes powerful; What is powerful becomes natural; What is natural becomes Tao.
What has Tao becomes everlasting And free from harm throughout life.
`
	passages[16] = `THE WAY OF SUBTLE INFLUENCE
Superior leaders are those whose existence is merely known;
The next best are loved and honored; The next are respected;
And the next are ridiculed.
Those who lack belief Will not in turn be believed. But when the command comes from afar And the work is done, the goal achieved, The people say, "We did it naturally."
`
	passages[17] = `LOSING THE INSTINCTS
When the great Tao is forgotten,
Philanthropy and morality appear. Intelligent strategies are produced,
And great hypocrisies emerge.
When the Family has no Harmony,
Piety and devotion appear. The nation is confused by chaos,
And loyal patriots emerge.
`
	passages[18] = `RETURN TO SIMPLICITY
Discard the sacred, abandon strategies;
The people will benefit a hundredfold. Discard philanthropy, abandon morality;
The people will return to natural love. Discard cleverness, abandon the acquisitive;
The thieves will exist no longer.
However, if these three passages are inadequate,
Adhere to these principles:
Perceive purity; Embrace simplicity; Reduce self-interest; Limit desires.
`
	passages[19] = `DEVELOPING INDEPENDENCE
Discard the academic; have no anxiety. How much difference is there between agreement and servility? How much difference is there between good and evil? That one should revere what others revere — how absurd and uncentered!
The Collective Mind is expansive and flourishing,
As if receiving a great sacrifice,
As if ascending a living observatory. I alone remain uncommitted,
Like an infant who has not yet smiled,
Unattached, without a place to merge. The Collective Mind is all-encompassing. I alone seem to be overlooked.
I am unknowing to the core and unclear, unclear!
Ordinary people are bright and obvious;
I alone am dark and obscure. Ordinary people are exacting and sharp;
I alone am subdued and dull.
Indifferent like the sea,
Ceaseless like a penetrating wind,
The Collective Mind is ever present And yet, I alone am unruly and remote.
I alone am different from the others
In treasuring nourishment from the Mother.
`
	passages[20] = `KNOWING THE COLLECTIVE ORIGIN
The natural expression of Power
Proceeds only through the Tao. The Tao acts through Natural Law;
So formless, so intangible.
Intangible, formless!
At its center appears the Image. Formless, intangible!
At its center appears Natural Law. Obscure, mysterious!
At its center appears the Life Force. The Life Force is very real;
At its center appears truth.
From ancient times to the present, Its name ever remains, Through the experience of the Collective Origin.
How do I know the way of the Collective Origin? Through this.
`
	passages[21] = `FOLLOWING THE PATTERN
What is curved becomes whole;
What is crooked becomes straight. What is deep becomes filled:
What is exhausted becomes refreshed. What is small becomes attainable;
What is excessive becomes confused.
Thus Evolved Individuals hold to the One And regard the world as their Pattern.
They do not display themselves;
Therefore they are illuminated. They do not define themselves;
Therefore they are distinguished. They do not make claims;
Therefore they are credited. They do not boast;
Therefore they advance.
Since, indeed, they do not compete, The world cannot compete with them.
That ancient saying: "What is curved becomes whole"
Are these empty words? To become whole,
Turn within.
`
	passages[22] = `THE STEADY FORCE OF ATTITUDE
Nature rarely speaks.
Hence the whirlwind does not last a whole morning,
Nor the sudden rainstorm last a whole day. What causes these? Heaven and Earth. If Heaven and Earth cannot make them long lasting, How much less so can humans?
Thus, those who cultivate the Tao
Identify with the Tao. Those who cultivate Power
Identify with Power. Those who cultivate failure
Identify with failure.
Those who identify with the Tao
Are likewise welcomed by the Tao. Those who identify with Power
Are likewise welcomed by Power. Those who identify with failure
Are likewise welcomed by failure.
Those who lack belief Will not in turn be believed.
`
	passages[23] = `THE DANGER IN EXCESS
Those who are on tiptoe cannot stand firm. Those who straddle cannot walk. Those who display themselves cannot illuminate. Those who define themselves cannot be distinguished. Those who make claims can have no credit. Those who boast cannot advance.
To those who stay with the Tao, These are like excess food and redundant actions And are contrary to Natural Law. Thus those who possess the Tao turn away.
`
	passages[24] = `THE TAO OF GREATNESS
There was something in a state of fusion Before Heaven and Earth were born.
Silent, vast,
Independent, and unchanging; Working everywhere, tirelessly;
It can be regarded as Mother of the world. I do not know its name;
The word I say is Tao. Forced to give it a name,
I say Great.
Great means continuing. Continuing means going far. Going far means returning.
Therefore the Tao is Great. Heaven and Earth are Great.
A leader is likewise Great. In the universe there are four Greatnesses, And leadership is one of them.
Humans are modeled on the earth.
The earth is modeled on heaven. Heaven is modeled on the Tao.
The Tao is modeled on nature.
`
	passages[25] = `THE GRAVITY OF POWER
Gravity is the foundation of levity. Stillness is the master of agitation.
Thus Evolved Individuals can travel the whole day
Without leaving behind their baggage. However arresting the views,
They remain calm and unattached. How can leaders with ten thousand chariots
Have a light-hearted position in the world?
If they are light-hearted, they lose their foundation. If they are agitated, they lose their mastery.
`
	passages[26] = `THE SKILLFUL EXCHANGE OF INFORMATION
A good path has no ruts. A good speech has no flaws. A good analysis uses no schemes.
A good lock has no bar or bolt,
And yet it cannot be opened. A good knot does not restrain,
And yet it cannot be unfastened.
Thus Evolved Individuals are always good at saving others;
Hence no one is wasted. They are always good at saving things;
Hence nothing is wasted.
This is called Doubling the Light.
Therefore a good person is the teacher of an inferior person;
And an inferior person is the resource of a good person. One who does not treasure a teacher, or does not cherish a resource,
Although intelligent, is greatly deluded.
This is called Significant Subtlety.
`
	passages[27] = `UNITING THE FORCES
Know the male,
Hold to the female;
Become the world's stream. By being the world's stream,
The Power will never leave.
This is returning to Infancy
Know the white,
Hold to the black;
Become the world's pattern. By becoming the world's pattern, The Power will never falter.
This is returning to Limitlessness.
Know the glory,
Hold to the obscurity;
Become the world's valley. By being the world's valley,
The Power will be sufficient.
This is returning to Simplicity.
When Simplicity is broken up, It is made into instruments. Evolved Individuals who employ them, Are made into leaders. In this way the Great System is united.
`
	passages[28] = ``
	passages[29] = ``
	passages[30] = ``
	passages[31] = ``
	passages[32] = ``
	passages[33] = ``
	passages[34] = ``
	passages[35] = ``
	passages[36] = ``
	passages[37] = ``
	passages[38] = ``
	passages[39] = ``
	passages[40] = ``
	passages[41] = ``
	passages[42] = ``
	passages[43] = ``
	passages[44] = ``
	passages[45] = ``
	passages[46] = ``
	passages[47] = ``
	passages[48] = ``
	passages[49] = ``
	passages[50] = ``
	passages[51] = ``
	passages[52] = ``
	passages[53] = ``
	passages[54] = ``
	passages[55] = ``
	passages[56] = ``
	passages[57] = ``
	passages[58] = ``
	passages[59] = ``
	passages[60] = ``
	passages[61] = ``
	passages[62] = ``
	passages[63] = ``
	passages[64] = ``
	passages[65] = ``
	passages[66] = ``
	passages[67] = ``
	passages[68] = ``
	passages[69] = ``
	passages[70] = ``
	passages[71] = ``
	passages[72] = ``
	passages[73] = ``
	passages[74] = ``
	passages[75] = ``
	passages[76] = ``
	passages[77] = `ACCEPTING THE BLAME
Nothing in the world,
  Is as yielding and receptive as water;
Yet in attacking the firm and inflexible,
  Nothing triumphs so well.
Because of what it is not,
  This becomes easy.

The receptive triumphs over the inflexible;
  The yielding triumphs over the rigid.
None in the world do not know this.
  None have the ability to practice it.

Therefore Evolved Individuals say:
  One who accepts the disgrace of the organization
  Can be called the leader of the grain shrine.
  One who accepts the misfortunes of the organization
  Can be called the leader of the world.

Right words appear to reverse themselves.
`
	passages[78] = `THE POWER IN NOT TAKING ADVANTAGE
Even when a great resentment is reconciled,
Some resentment must linger.
How can this be made good?

That is why Evolved Individuals
Hold the left side of the contract
And do not censure others.
Those with Power are in charge of the contract;
Those without Power are in charge of resolving it.

The Tao in Nature has no favorites.
It always works through the good person.
`
	passages[79] = `FULFILLING INDEPENDENCE
In a small organization with few people;

Let there be ten or a hundred times
  More tools than they can use.
Let the people value their lives
  And yet not move far away.
Even though there are boats and carriages,
  There is no occasion to use them.
Even though there are armor and weapons,
  There is no occasion to display them.

Let the people again knot cords and
  Their food will be pleasing.
  Their clothes will be fine.
  Their homes will be secure.
  Their customs will be joyful.

Nearby organizations may watch each other;
  Their crowing and barking may be heard.
Yet the people may grow old and die
  Without coming or going between them.
`
	passages[80] = `THE EVOLVED WAY
Sincere words are not embellished;
  Embellished words are not sincere.
Those who are good are not defensive;
  Those who are defensive are not good.
Those who know are not erudite;
  Those who are erudite do not know.

Evolved Individuals do not accumulate.
  The more they do for others, the more they gain;
  The more they give to others, the more they possess.

The Tao of Nature
  Is to serve without spoiling.
The Tao of Evolved Individuals
  Is to act without contending.
`

	index := rand.Intn(80)
	return index, passages[index]
}
