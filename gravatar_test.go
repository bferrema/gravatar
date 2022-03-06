package gravatar

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want *Gravatar
	}{
		{
			name: "default",
			args: args{
				email: "brunom4ciel@gmail.com",
			},
			want: &Gravatar{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.email); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGravatar_URL(t *testing.T) {
	type fields struct {
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default",
			fields: fields{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
			},
			want: "https://www.gravatar.com/55c9274ab3925685a90dd838f0eaccd9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gravatar{
				hash:         tt.fields.hash,
				defaultURL:   tt.fields.defaultURL,
				defaultValue: tt.fields.defaultValue,
				size:         tt.fields.size,
				forceDefault: tt.fields.forceDefault,
				rating:       tt.fields.rating,
			}
			if got := g.URL(); got != tt.want {
				t.Errorf("Gravatar.URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGravatar_AvatarURL(t *testing.T) {
	type fields struct {
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default",
			fields: fields{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9",
		},
		{
			name: "defaultValue",
			fields: fields{
				hash:       "55c9274ab3925685a90dd838f0eaccd9",
				defaultURL: string(NotFound),
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?d=404",
		},
		{
			name: "size",
			fields: fields{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
				size: 200,
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?s=200",
		},
		{
			name: "forceDefault",
			fields: fields{
				hash:         "55c9274ab3925685a90dd838f0eaccd9",
				forceDefault: true,
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?f=y",
		},
		{
			name: "rating",
			fields: fields{
				hash:   "55c9274ab3925685a90dd838f0eaccd9",
				rating: string(Pg),
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?r=pg",
		},
		{
			name: "combined",
			fields: fields{
				hash:       "55c9274ab3925685a90dd838f0eaccd9",
				defaultURL: string(NotFound),
				size:       200,
			},
			want: "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?d=404&s=200",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gravatar{
				hash:         tt.fields.hash,
				defaultURL:   tt.fields.defaultURL,
				defaultValue: tt.fields.defaultValue,
				size:         tt.fields.size,
				forceDefault: tt.fields.forceDefault,
				rating:       tt.fields.rating,
			}
			if got := g.AvatarURL(); got != tt.want {
				t.Errorf("Gravatar.AvatarURL() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("defaultURL", func(t *testing.T) {
		g := &Gravatar{
			hash: "55c9274ab3925685a90dd838f0eaccd9",
		}

		g.DefaultURL("https://example.com/escaping test")
		want := "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9?d=https://example.com/escaping%20test"
		if got := g.AvatarURL(); got != want {
			t.Errorf("Gravatar.AvatarURL() = %v, want %v", got, want)
		}
	})

	t.Run("invalidSize", func(t *testing.T) {
		g := &Gravatar{
			hash: "55c9274ab3925685a90dd838f0eaccd9",
		}

		g.Size(-1)
		want := "https://www.gravatar.com/avatar/55c9274ab3925685a90dd838f0eaccd9"
		if got := g.AvatarURL(); got != want {
			t.Errorf("Gravatar.AvatarURL() = %v, want %v", got, want)
		}
	})
}

func TestGravatar_JSONURL(t *testing.T) {
	type fields struct {
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "default",
			fields: fields{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
			},
			want: "https://www.gravatar.com/55c9274ab3925685a90dd838f0eaccd9.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gravatar{
				hash:         tt.fields.hash,
				defaultURL:   tt.fields.defaultURL,
				defaultValue: tt.fields.defaultValue,
				size:         tt.fields.size,
				forceDefault: tt.fields.forceDefault,
				rating:       tt.fields.rating,
			}
			if got := g.JSONURL(); got != tt.want {
				t.Errorf("Gravatar.JSONURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGravatar_JSONURLCallback(t *testing.T) {
	type fields struct {
		hash         string
		defaultURL   string
		defaultValue string
		size         int
		forceDefault bool
		rating       string
	}
	type args struct {
		callback string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "default",
			fields: fields{
				hash: "55c9274ab3925685a90dd838f0eaccd9",
			},
			args: args{
				callback: "alert",
			},
			want: "https://www.gravatar.com/55c9274ab3925685a90dd838f0eaccd9.json?callback=alert",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gravatar{
				hash:         tt.fields.hash,
				defaultURL:   tt.fields.defaultURL,
				defaultValue: tt.fields.defaultValue,
				size:         tt.fields.size,
				forceDefault: tt.fields.forceDefault,
				rating:       tt.fields.rating,
			}
			if got := g.JSONURLCallback(tt.args.callback); got != tt.want {
				t.Errorf("Gravatar.JSONURLCallback() = %v, want %v", got, tt.want)
			}
		})
	}
}
