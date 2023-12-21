from django.utils import timezone
import uuid
from django.db import models

# Create your models here.
class League(models.Model):
    name = models.CharField(max_length=250, null=False, blank=False, unique=False)
    description = models.TextField(null=True, blank=True)

    date_created = models.DateTimeField(default=timezone.now)
    date_updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.name

class Team(models.Model):
    name = models.CharField(max_length=250, null=False, blank=False, unique=False)
    description = models.TextField(null=True, blank=True)

    date_created = models.DateTimeField(default=timezone.now)
    date_updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.name

class Player(models.Model):
    id = models.UUIDField(default=uuid.uuid4, primary_key=True, editable=False, serialize=False)
    name = models.CharField(max_length=250, null=False, blank=False)
    team = models.ForeignKey(Team, on_delete=models.CASCADE, related_name='players')

    date_created = models.DateTimeField(default=timezone.now)
    date_updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.name

class Match(models.Model):
    id = models.UUIDField(default=uuid.uuid4, primary_key=True, editable=False, serialize=False)
    home_team = models.ForeignKey(Team, on_delete=models.CASCADE, related_name='home_matches')
    away_team = models.ForeignKey(Team, on_delete=models.CASCADE, related_name='away_matches')
    home_team_score = models.IntegerField(null=False, blank=False)
    away_team_score = models.IntegerField(null=False, blank=False)
    date = models.DateTimeField(null=False, blank=False)
    league = models.ForeignKey(League, on_delete=models.CASCADE, related_name='matches')

    date_created = models.DateTimeField(default=timezone.now)
    date_updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return f'{self.home_team.name} vs {self.away_team.name}'

class Spectator(models.Model):
    id = models.UUIDField(default=uuid.uuid4, primary_key=True, editable=False, serialize=False)
    name = models.CharField(max_length=250, null=False, blank=False)
    match = models.ForeignKey(Match, on_delete=models.CASCADE, related_name='spectators')

    date_created = models.DateTimeField(default=timezone.now)
    date_updated = models.DateTimeField(auto_now=True)

    def __str__(self):
        return self.name