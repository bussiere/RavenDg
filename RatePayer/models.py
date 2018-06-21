from django.db import models

# Create your models here.

class TaxUser(models.Model):
    name = models.CharField(blank=True,null=True,max_length=50)
    surname = models.CharField(blank=True, null=True, max_length=50)
    spi =  models.CharField(blank=True, null=True, max_length=50)
    email =  models.CharField(blank=True, null=True, max_length=50)
    note = models.CharField(blank=True, null=True, max_length=200)
